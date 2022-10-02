package main

import (
	"flag"

	"apihut-server/config"
	"apihut-server/dao/mysql"
	"apihut-server/dao/redis"
	"apihut-server/logger"
	"apihut-server/routers"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "f", "./conf/config.yaml", "Config file")
}

func main() {
	flag.Parse()

	var err error
	// 初始化配置
	err = config.Init(configFile)
	if err != nil {
		panic(err)
	}
	// 初始化日志
	err = logger.Init()
	if err != nil {
		panic(err)
	}
	// 初始化路由
	r := routers.SetupRouter()
	// 初始化数据库
	err = mysql.Init()
	if err != nil {
		panic(err)
	}
	// 初始化Redis
	err = redis.Init()
	if err != nil {
		panic(err)
	}

	_ = r.Run(config.GetSitePort())
}
