package bleve

import (
	"encoding/json"
	"strconv"

	"github.com/apihutco/server/dao/db"
	"github.com/apihutco/server/logger"
	"github.com/apihutco/server/utils/gen"

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
	//if err = SyncFromDB(); err != nil {
	//	return err
	//}

	return nil
}

var lastSyncMD5 string

func SyncFromDB() error {
	greetList, err := db.Ctrl().Greet().List()
	if err != nil {
		return err
	}
	if len(greetList) == 0 {
		return ErrorNotFound
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
