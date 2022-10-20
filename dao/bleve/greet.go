package bleve

import (
	"apihut-server/models"
	"github.com/blevesearch/bleve/v2"
	uquery "github.com/blevesearch/bleve/v2/search/query"
	"github.com/mitchellh/mapstructure"
)

func SearchGreet(str string) ([]*models.Greet, error) {
	query := bleve.NewMatchQuery(str)
	query.SetOperator(uquery.MatchQueryOperatorAnd)
	search := bleve.NewSearchRequest(query)
	// search.Size = 1
	// search.Fields = []string{"sentence", "author", "tags"}
	search.Fields = []string{"sentence"}
	rsp, err := i.greet.Search(search)
	if err != nil {
		return nil, err
	}

	re := make([]*models.Greet, 0)
	for _, document := range rsp.Hits {
		item := new(models.Greet)
		if err = mapstructure.Decode(document.Fields, &item); err != nil {
			return nil, err
		}
		re = append(re, item)
	}

	return re, nil
}