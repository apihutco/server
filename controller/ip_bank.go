package controller

import (
	"apihut-server/logger"
	"apihut-server/logic/ip_bank"
	"apihut-server/response"
	"net"

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
		response.BadRequest(c).Code(response.ErrorFormat).JSON()
		return
	}

	info, err := ip_bank.GetIP(ip)
	if err != nil {
		logger.L().Error("无法定位", zap.Error(err), zap.String("IP", ip.String()))
		response.Error(c).Code(response.ErrorIPUnableToLocate).Data(gin.H{"ip": ip.String()}).JSON()
		return
	}

	response.Success(c).Data(info).JSON()
}

func IPTextHandler(c *gin.Context) {
	response.Success(c).Data(c.ClientIP() + "\n").String()
}
