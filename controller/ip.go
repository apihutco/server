package controller

import (
	"apihut-server/logger"
	"apihut-server/logic"
	"apihut-server/response"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net"
)

func IPHandler(c *gin.Context) {
	strIP := c.Param("ip")
	if len(strIP) == 0 {
		strIP = c.ClientIP()
	}

	ip := net.ParseIP(strIP)
	info, err := logic.GetIP(ip)
	if err != nil {
		logger.L().Error("IP not found", zap.Error(err))
		response.ErrorWithMsg(c, errors.Unwrap(err).Error())
		return
	}

	response.SuccessWithData(c, info)
}
