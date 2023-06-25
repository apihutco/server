package routers

import (
	"github.com/apihutco/server/config"
	. "github.com/apihutco/server/controller"
	"github.com/apihutco/server/routers/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.Conf.Mode)

	r := gin.New()

	// 使用 CloudFlare Tunnel 时获取客户端真实地址
	r.TrustedPlatform = gin.PlatformCloudflare

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	r.StaticFile("favicon.ico", "./static/favicon.ico")
	r.Use(middleware.Logger())

	// IP定位
	ip := r.Group("/ip")
	{
		// JSON
		ip.GET("", IPHandler) // 请求来源IP
		ip.GET("/:ip", IPHandler)
	}

	// 协议测试（get，post，ws）
	{
		r.GET("", GetHandler)                   // JSON形式返回Query参数，为空即为首页
		r.POST("", PostHandler)                 // 原样返回请求的Body
		r.GET("/ws", WebSocketHandler)          // 单机ws收发
		r.GET("/ws/:channel", WebSocketHandler) // 频道ws收发
	}

	// 哈希头像生成
	{
		r.GET("/avatar", AvatarHandler)
		r.GET("/avatar/:hash", AvatarHandler)
	}

	// 一句招呼
	{
		r.GET("/greet", GreetHandler)
	}

	// 地理位置
	{
		r.GET("/geo", GeoBankHandler)
	}

	// 天气
	{
		w := r.Group("/weather")
		w.GET("/now", WeatherNowHandler)
		w.GET("/3day", WeatherDay3Handler)
	}

	// 健康检查
	r.GET("/health", HealthHandler)
	// 无匹配
	r.NoRoute(NotFound)

	return r
}
