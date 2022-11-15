package greet

import (
	"strings"

	"apihut-server/dao/bleve"
	"apihut-server/logger"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"apihut-server/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func GetGreet(s string) (*models.Greet, error) {
	var query string
	query = s
	if len(query) == 0 {
		q := strings.Builder{}
		q.WriteString(consts.GetTimeCode().String())
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
