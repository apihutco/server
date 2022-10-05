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
	ip := r.Group("/ip")
	{
		// JSON
		ip.GET("", IPJSONHandler)          // 请求来源IP
		ip.GET("/:ip", IPJSONHandler)      // 指定IP
		ip.GET("/json/:ip", IPJSONHandler) // JSON形式完整版
		// 纯文字
		ip.GET("/text", IPTextHandler) // 纯文字形式返回
	}

	// 协议测试（get，post，ws）
	r.GET("/get", GetHandler)         // JSON形式返回Query参数
	r.GET("/get/:output", GetHandler) // 按格式返回Query参数
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

	// 健康检查
	r.GET("/health", HealthHandler)

	// 无匹配
	r.NoRoute(NotFound)

	return r
}
