package geo_bank

import (
	"errors"
	"time"

	"apihut-server/config"
	"apihut-server/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var geoCtrlList []IGeoCtrl

type IGeoCtrl interface {
	GetInfo(location string, adm string) (*models.GeoBank, error)
	SaveInfo(info *models.GeoBank) error
}

func init() {
	go func() {
		time.Sleep(time.Second * 3)
		InitGeoCtrl(config.Conf.QWeather.Key)
		info, err := GetGeoInfo("深圳", "")
		if err != nil {
			zap.L().Error("get geo info ", zap.Error(err))
			return
		}
		zap.L().Info("info", zap.Any("info", info))
	}()
}

func InitGeoCtrl(key string) {
	geoCtrlList = make([]IGeoCtrl, 0)
	geoCtrlList = append(geoCtrlList,
		NewLocal(),
		NewQWeather(key),
	)
}

func GetGeoInfo(location, adm string) (*models.GeoBank, error) {
	// 按序遍历
	for _, ctrl := range geoCtrlList {
		info, err := ctrl.GetInfo(location, adm)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			zap.L().Error("Get Geo", zap.Error(err))
		}

		if err == nil && info != nil {
			err = ctrl.SaveInfo(info)
			if err != nil {
				zap.L().Error("Get Info Save", zap.Error(err))
			}
			return info, nil
		}
	}

	return nil, errors.New("Geo Info Not Found")
}
