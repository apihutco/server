package controller

import (
	"apihut-server/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func HomeHandler(c *gin.Context) {
	logger.L().Info("home handler", zap.String("time", time.Now().String()))
}
