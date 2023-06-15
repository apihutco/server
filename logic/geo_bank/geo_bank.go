package geo_bank

import (
	"errors"

	"apihut-server/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var geoCtrlList []IGeoCtrl

type IGeoCtrl interface {
	GetInfo(location string, adm string) (*models.GeoBank, error)
	SaveInfo(info *models.GeoBank) error
}

func InitGeoCtrl() {
	geoCtrlList = make([]IGeoCtrl, 0)
	geoCtrlList = append(geoCtrlList,
		NewLocal(),
		NewQWeather(),
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
