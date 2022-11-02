package controller

import (
	"apihut-server/logger"
	"apihut-server/utils/ws"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"

	"apihut-server/logic/consts"
	"apihut-server/logic/protocol"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {

	output := c.Param("output")

	switch output {
	case consts.CaseOutputType(output, consts.Text):
		str := protocol.ParamsToString(c.Request)
		c.String(http.StatusOK, str)
	default:
		res := protocol.ParamsToJSON(c.Request)
		response.SuccessWithData(c, res)
	}
}

func PostHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	if err != nil {
		logger.L().Error("读取Body失败", zap.Error(err))
		response.ErrorWithCode(c, response.ErrorProtocolReadBody)
		return
	}

	var h gin.H
	err = json.Unmarshal(body, &h)
	if err != nil {
		logger.L().Error("序列化失败", zap.Error(err))
		response.ErrorWithCode(c, response.ErrorProtocolUnmarshal)
		return
	}

	response.SuccessWithData(c, h)
}

var hub *ws.Hub

func init() {
	hub = ws.NewHub()
	go hub.Run()
}

func WebSocketHandler(c *gin.Context) {
	err := ws.Handler(hub, c)
	if err != nil {
		logger.L().Error("协议升级失败", zap.Error(err))
		response.ErrorWithCode(c, response.ErrorProtocolWsUpgrade)
		return
	}
}
