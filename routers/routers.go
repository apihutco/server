package routers

import (
	"apihut-server/config"
	. "apihut-server/controller"
	"apihut-server/utils/ws"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.Share.Mode)

	r := gin.New()
	r.StaticFile("favicon.ico", "./assets/favicon.ico")

	// 首页
	r.GET("/", HomeHandler)

	// IP定位
	r.GET("/ip", IPHandler)     // 请求来源IP
	r.GET("/ip/:ip", IPHandler) // 指定IP

	// 协议测试（get，post，ws）
	r.GET("/get", GetHandler)
	r.POST("/post", PostHandler)

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		ws.Handler(hub, c)
	})
	r.GET("/ws/:channel", func(c *gin.Context) {
		ws.Handler(hub, c)
	})

	// 哈希头像生成
	r.GET("/avatar", AvatarHandler)
	r.GET("/avatar/:hash", AvatarHandler)

	// 网课题库

	// 健康检查
	r.GET("/health", HealthHandler)

	return r
}
