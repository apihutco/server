package bleve

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"

	"apihut-server/config"
	"apihut-server/dao/mysql"
	"apihut-server/logger"

	"github.com/blevesearch/bleve/v2"
	"github.com/pkg/errors"
	gse "github.com/vcaesar/gse-bleve"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

	// 同步索引
	if err = migrate(); err != nil {
		return err
	}

	return nil
}

func migrate() error {
	greetList, err := mysql.GetGreetList()
	if errors.Is(gorm.ErrRecordNotFound, err) || len(greetList) == 0 {
		if err = loadSQLFile(); err != nil {
			return err
		}
		greetList, err = mysql.GetGreetList()
	} else if err != nil {
		return err
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
