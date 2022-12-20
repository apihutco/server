package cron

import (
	"apihut-server/config"
	"apihut-server/dao/bleve"
	"apihut-server/logger"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Init() {
	c := cron.New()

	// 一句招呼数据同步
	_, _ = c.AddFunc(config.Conf.Bleve.SyncCron, func() {
		logger.L().Info("【定时任务】开始同步一句招呼数据...")
		err := bleve.SyncFromDB()
		if err != nil {
			logger.L().Error("【定时任务】同步一句招呼数据失败", zap.Error(err))
			return
		}
		logger.L().Info("【定时任务】同步一句招呼数据成功！")
	})

	c.Start()
}
