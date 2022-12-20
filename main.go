package main

import (
	"flag"

	"apihut-server/dao/bleve"

	"apihut-server/config"
	"apihut-server/dao/mysql"
	"apihut-server/dao/redis"
	"apihut-server/logger"
	"apihut-server/routers"

	"go.uber.org/zap"
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
		logger.L().DPanic("database panic", zap.Error(err))
		return
	}
	// 初始化Redis
	err = redis.Init()
	if err != nil {
		logger.L().DPanic("redis panic", zap.Error(err))
		return
	}
	// 初始化全文索引
	err = bleve.Init(config.Conf.Bleve.Index)
	if err != nil {
		logger.L().DPanic("bleve panic", zap.Error(err))
		return
	}

	_ = r.Run(config.Conf.GetSitePort())
}
