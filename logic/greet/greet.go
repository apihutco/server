package greet

import (
	"strings"

	"github.com/apihutco/server/dao/bleve"
	"github.com/apihutco/server/logger"
	"github.com/apihutco/server/models"
	"github.com/apihutco/server/utils"
	"github.com/apihutco/server/utils/consts"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func GetGreet(s string) (*models.Greet, error) {
	var query string
	query = s
	if len(query) == 0 {
		q := strings.Builder{}
		q.WriteString(consts.RepoTime.Now().String())
		query = q.String()
	}

	logger.L().Debug("一句招呼查询", zap.String("条件", query))

	ls, err := bleve.SearchGreet(query)
	if err != nil && !errors.Is(err, bleve.ErrorNotFound) {
		return nil, err
	}

	if errors.Is(err, bleve.ErrorNotFound) {
		ls, err = bleve.SearchGreet(consts.DefaultCode.CN())
	}

	logger.L().Debug("一句招呼查询", zap.Any("结果", ls))

	if err != nil {
		return nil, err
	}

	return ls[utils.GetRange(len(ls))], nil
}
