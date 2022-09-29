package routers

import (
	"apihut-server/config"
	. "apihut-server/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.Share.Mode)

	r := gin.New()
	// 首页
	r.GET("/", HomeHandler)
	// 文档

	// IP定位（腾讯，高德）
	r.GET("/ip", IPHandler)     // 请求来源IP
	r.GET("/ip/:ip", IPHandler) // 指定IP

	// 协议测试（get，post，ws）

	// 哈希头像生成（自有api风格，gravatar风格，https://www.gravatar.com/avatar/HASH）
	// https://en.gravatar.com/site/implement/images/
	// 支持设置默认头像

	// 网课题库

	// 健康检查
	r.GET("/health", HealthHandler)

	return r
}
