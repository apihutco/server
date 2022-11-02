package controller

import (
	"apihut-server/logger"
	"apihut-server/logic/greet"
	"apihut-server/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GreetHandler(c *gin.Context) {

	str := c.Query("s")

	re, err := greet.GetGreet(str)
	if err != nil {
		logger.L().Error("获取一句招呼失败", zap.Error(err), zap.String("query", str))
		response.Error(c).JSON()
		return
	}

	response.Success(c).Data(re).JSON()
}
