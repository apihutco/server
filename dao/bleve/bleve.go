package bleve

import (
	"apihut-server/dao/mysql"
	"github.com/blevesearch/bleve/v2"
	"github.com/pkg/errors"
	gse "github.com/vcaesar/gse-bleve"
	"strconv"
)

var i *index

type index struct {
	greet bleve.Index
}

func Init(indexPath string) error {
	var greetIndex bleve.Index
	var err error

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
	// if err = migrate(); err != nil {
	// 	return err
	// }

	return nil
}

func migrate() error {
	greetList, err := mysql.GetGreetList()
	if err != nil {
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
