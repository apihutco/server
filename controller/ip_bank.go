package controller

import (
	"net"

	"apihut-server/logger"
	"apihut-server/logic/ip_bank"
	"apihut-server/response"
	"apihut-server/utils/consts"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func IPHandler(c *gin.Context) {
	strIP := c.Param("ip")
	if len(strIP) == 0 {
		strIP = c.ClientIP()
	}

	strOutput := c.Query("output")
	if len(strOutput) == 0 {
		strOutput = c.DefaultQuery("o", consts.RepoOutput.JSON.String())
	}

	logger.L().Debug("request", zap.String("ip", strIP), zap.String("output", strOutput))

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

	info.CacheTime = info.UpdatedAt

	// 按格式返回
	switch true {
	case consts.RepoOutput.Is(strOutput, consts.RepoOutput.Text):
		response.Success(c).Data(info.String()).String()
	default:
		response.Success(c).Data(info).JSON()
	}
}
