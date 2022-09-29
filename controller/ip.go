package controller

import (
	"apihut-server/logger"
	"apihut-server/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func IPHandler(c *gin.Context) {
	strIP := c.Param("ip")
	if len(strIP) == 0 {
		strIP = c.ClientIP()
	}

	ip := net.ParseIP(strIP)
	info, err := logic.GetIP(ip)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		logger.L().Error("IP info not found", zap.Error(err))
		return
	}

	c.JSON(200, info)
}
