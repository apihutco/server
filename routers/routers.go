package routers

import (
	"apihut-server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.ShareConf.Mode)

	r := gin.New()
	// 首页
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	// 文档

	// IP定位（腾讯，高德）

	// 协议测试（get，post，ws）

	// 哈希头像生成（自有api风格，gravatar风格，https://www.gravatar.com/avatar/HASH）
	// https://en.gravatar.com/site/implement/images/
	// 支持设置默认头像

	// 网课题库

	return r
}
