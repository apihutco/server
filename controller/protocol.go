package controller

import (
	"apihut-server/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func GetHandler(c *gin.Context) {
	response.SuccessWithData(c, gin.H{"hello": "world"})
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
