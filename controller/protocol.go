package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/apihutco/server/logger"
	"github.com/apihutco/server/logic/home"
	"github.com/apihutco/server/logic/protocol"
	"github.com/apihutco/server/response"
	"github.com/apihutco/server/utils/consts"
	"github.com/apihutco/server/utils/ws"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHandler(c *gin.Context) {

	if len(c.Request.URL.Query()) == 0 {
		page := home.Page()
		response.Success(c).Data(page).String()
		return
	}

	output := c.Query("output")
	if len(output) == 0 {
		output = c.DefaultQuery("o", consts.RepoOutput.JSON.String())
	}

	switch true {
	case consts.RepoOutput.Is(output, consts.RepoOutput.Text):
		str := protocol.ParamsToString(c.Request)
		response.Success(c).Data(str).String()
	default:
		res := protocol.ParamsToJSON(c.Request)
		response.Success(c).Data(res).Pure()
	}
}

func PostHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	if err != nil {
		logger.L().Error("读取Body失败", zap.Error(err))
		response.BadRequest(c).Code(response.ErrorProtocolReadBody).JSON()
		return
	}

	switch c.Request.Header.Get("Content-Type") {
	case "application/json":
		var h gin.H
		err = json.Unmarshal(body, &h)
		if err != nil {
			logger.L().Error("序列化失败", zap.Error(err))
			response.Error(c).Code(response.ErrorProtocolUnmarshal).JSON()
			return
		}
		response.Success(c).Data(h).Pure()
	default:
		response.Success(c).Data(string(body)).Pure()
	}

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
		response.Error(c).Code(response.ErrorProtocolWsUpgrade).JSON()
		return
	}
}
