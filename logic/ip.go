package logic

import (
	"apihut-server/config"
	"apihut-server/dao/mysql"
	"apihut-server/logger"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net"
)

type IIPCtrl interface {
	GetIP(net.IP) (*models.IPBank, error)
	Platform() consts.PlatformID
}

func GetIP(ip net.IP) (*models.IPBank, error) {
	ctrlList := make([]IIPCtrl, 3)
	ctrlList[0] = LocalInit()
	ctrlList[1] = TencentInit(config.Share.Open.Tencent.Key)
	// ctrlList[2] = GaodeInit(config.Share.Open.Gaode.Key)

	var allErr error
	// 按序轮询数据源
	for i := 0; i < len(ctrlList); i++ {
		if ctrlList[i] == nil {
			continue
		}
		ctrl := ctrlList[i]
		ipInfo, err := ctrl.GetIP(ip)
		if err != nil {
			logger.L().Error("Get IP", zap.Error(err), zap.String("platform", ctrl.Platform().Name()))
			allErr = errors.WithMessagef(err, "Platform: %s,Err", ctrl.Platform().Name())
			continue
		}
		// 非数据库来源的，持久化到数据库
		if i > 0 && ipInfo != nil {
			err = mysql.CreateIPBank(ipInfo)
			if err != nil {
				logger.L().Error("Save to db", zap.Error(err), zap.String("from", ctrl.Platform().Name()), zap.Any("info", ipInfo))
			}
			// 持久化失败也返回
			return ipInfo, errors.WithMessagef(err, "Platform: %s,Info: %+v,Err", ctrl.Platform().Name(), ipInfo)
		}
	}
	return nil, allErr
}
