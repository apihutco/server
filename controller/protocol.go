package controller

import (
	"encoding/json"
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
		response.Error(c)
		return
	}

	var h gin.H
	_ = json.Unmarshal(body, &h)

	response.SuccessWithData(c, h)
}

func WebSocketHandler(c *gin.Context) {

}

func WebSocketWithChannel(c *gin.Context) {

}
