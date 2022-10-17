package controller

import (
	"net"
	"net/http"

	"apihut-server/logger"
	"apihut-server/logic/ip_bank"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func IPJSONHandler(c *gin.Context) {
	strIP := c.Param("ip")
	if len(strIP) == 0 {
		strIP = c.ClientIP()
	}

	ip := net.ParseIP(strIP)
	if ip == nil {
		response.ErrorWithMsg(c, "IP格式错误")
		return
	}

	info, err := ip_bank.GetIP(ip)
	if err != nil {
		logger.L().Error("IP not found", zap.Error(err))
		response.ErrorWithMsgAndData(c, "无法定位", gin.H{"ip": strIP})
		return
	}

	response.SuccessWithData(c, info)
}

func IPTextHandler(c *gin.Context) {
	c.String(http.StatusOK, c.ClientIP())
}
