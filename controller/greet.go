package controller

import (
	"github.com/apihutco/server/logger"
	"github.com/apihutco/server/logic/greet"
	"github.com/apihutco/server/response"
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
