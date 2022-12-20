package bleve

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"

	"apihut-server/dao/mysql"
	"apihut-server/logger"
	"apihut-server/utils/gen"

	"github.com/blevesearch/bleve/v2"
	"github.com/pkg/errors"
	gse "github.com/vcaesar/gse-bleve"
	"go.uber.org/zap"
)

var i *index

type index struct {
	greet bleve.Index
}

func Init(indexPath string) error {
	var greetIndex bleve.Index
	var err error

	if len(indexPath) == 0 {
		return ErrorIndexEmpty
	}

	opt := gse.Option{
		Index: indexPath,
		Dicts: "zh_s,zh_t,embed,js",
		Stop:  "",
		Opt:   "search-hmm",
		Trim:  "trim",
	}
	greetIndex, err = gse.New(opt)
	if err != nil {
		if errors.Is(err, bleve.ErrorIndexPathExists) {
			greetIndex, err = bleve.Open(indexPath)
		} else {
			return err
		}
	}

	i = new(index)
	i.greet = greetIndex

	// 同步索引，只增改不删除
	if err = SyncFromDB(); err != nil {
		return err
	}

	return nil
}

var lastSyncMD5 string

func SyncFromDB() error {
	greetList, err := mysql.GetGreetList()
	if errors.Is(gorm.ErrRecordNotFound, err) || len(greetList) == 0 {
		if err = loadSQLFile(); err != nil {
			return err
		}
		greetList, err = mysql.GetGreetList()
	} else if err != nil {
		return err
	}

	// 缓存上次变动的md5，避免频繁更新
	m, _ := json.Marshal(greetList)
	newMD5 := gen.MD5(m)
	if lastSyncMD5 == newMD5 {
		logger.L().Debug("md5相同，数据无变动", zap.String("md5", newMD5))
		return nil
	} else {
		logger.L().Debug("md5变动，数据更新", zap.String("old.md5", lastSyncMD5), zap.String("new.md5", newMD5))
		lastSyncMD5 = newMD5
	}

	batch := i.greet.NewBatch()
	for _, greet := range greetList {
		err = batch.Index(strconv.Itoa(int(greet.ID)), greet)
		if err != nil {
			return err
		}
	}

	if err = i.greet.Batch(batch); err != nil {
		return err
	}

	return nil
}

func loadSQLFile() error {
	fileList, err := os.ReadDir(config.Share.Bleve.SetupPath)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	for _, entry := range fileList {
		wg.Add(1)
		go func(entry os.DirEntry) {
			file, err := os.Open(path.Join(config.Share.Bleve.SetupPath, entry.Name()))
			defer func() {
				_ = file.Close()
				wg.Done()
			}()
			if err != nil {
				logger.L().Error("Setup open file", zap.Error(err))
				return
			}
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				logger.L().Error("Setup read all", zap.Error(err))
				return
			}

			logger.L().Debug("Setup SQL", zap.ByteString("content", bytes))

			err = mysql.Exec(string(bytes))
			if err != nil {
				logger.L().Error("Setup sql exec", zap.Error(err))
				return
			}
		}(entry)
	}

	wg.Wait()

	return nil
}
