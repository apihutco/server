package ip_bank

import (
	"apihut-server/dao/mysql"
	"apihut-server/logger"
	"apihut-server/logic/consts"
	"apihut-server/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net"
)

type IIPCtrl interface {
	GetIP(net.IP) (*models.IPBank, error)
	Platform() consts.PlatformID
}

func GetIP(ip net.IP) (*models.IPBank, error) {
	ctrlList := make([]IIPCtrl, 3)
	ctrlList[0] = LocalInit()
	ctrlList[1] = GaodeInit()
	ctrlList[2] = TencentInit()

	// 按序轮询数据源
	for i := 0; i < len(ctrlList); i++ {
		if ctrlList[i] == nil {
			continue
		}
		ctrl := ctrlList[i]
		ipInfo, err := ctrl.GetIP(ip)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			logger.L().Error("Get IP", zap.Error(err), zap.String("platform", ctrl.Platform().Name()))
			return nil, err
		}

		if ipInfo != nil {
			// 非数据库来源的，持久化到数据库
			var saveErr error
			if i != 0 {
				saveErr = mysql.CreateIPBank(ipInfo)
				if saveErr != nil {
					logger.L().Error("Save to db", zap.Error(saveErr), zap.String("from", ctrl.Platform().Name()), zap.Any("info", ipInfo))
				}
			}
			// 持久化成功与否都返回
			return ipInfo, saveErr
		}
	}
	return nil, errors.New("IP Not Found")
}
