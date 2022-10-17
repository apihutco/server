package controller

import (
	"apihut-server/dao/bleve"
	"apihut-server/logger"
	"apihut-server/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GreetHandler(c *gin.Context) {

	str := c.Query("s")

	re, err := bleve.SearchGreet(str)
	if err != nil {
		logger.L().Error("获取一句招呼失败", zap.Error(err))
		response.Error(c)
		return
	}

	response.SuccessWithData(c, re)
}
