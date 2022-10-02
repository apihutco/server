package controller

import (
	"strconv"
	"time"

	"apihut-server/logic/avatar"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"apihut-server/response"

	"github.com/gin-gonic/gin"
)

/*
参考文档：
https://en.gravatar.com/site/implement/images/
*/

func AvatarHandler(c *gin.Context) {
	hash := c.Param("hash")
	if len(hash) == 0 {
		hash = c.ClientIP()
	}

	req := models.NewAvatar(hash)

	// 只要出现random参数，不管值是什么，都开启随机模式
	if _, has := c.GetQuery("random"); has {
		hash = strconv.Itoa(int(time.Now().UnixNano()))
		req.Hash = hash
		req.Random = true
	}

	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}

	file, err := avatar.NewAvatar(&req)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	// response.Success(c)

	// 按格式返回
	if req.GetOutput() == consts.JSON {
		response.Success(c)
		return
	} else {
		c.File(file)
		return
	}
}
