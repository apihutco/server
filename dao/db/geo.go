package db

import (
	"sync"

	"github.com/apihutco/server/models"
	"gorm.io/gorm"
)

type IGeoBank interface {
	Create(*models.GeoBank) error
	Get(location string, adm string) (*models.GeoBank, error)
}

type geoCtrl struct {
	db *gorm.DB
}

var shareGeoBankCtrl *geoCtrl
var geoBankCtrlOnce sync.Once

func (data *Database) Geo() IGeoBank {
	geoBankCtrlOnce.Do(func() {
		shareGeoBankCtrl = &geoCtrl{
			db: data.db,
		}
	})

	return shareGeoBankCtrl
}

func (geo *geoCtrl) Create(geoInfo *models.GeoBank) error {
	return geo.db.Create(&geoInfo).Error
}

func (geo *geoCtrl) Get(location string, adm string) (*models.GeoBank, error) {
	info := new(models.GeoBank)
	var err error

	if adm != "" {
		sql := "(district LIKE ? or city LIKE ?) and (city LIKE ? or province LIKE ?)"
		params := []any{location + "%", location + "%", adm + "%", adm + "%"}
		err = geo.db.Where(sql, params...).First(&info).Error
		//if err == nil {
		//	return info, err
		//}
		//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//	return nil, err
		//}

		return info, err

		// 降级为根据 adm 查
		//sql = "city LIKE ? or province LIKE ?"
		//params = []any{adm + "%", adm + "%"}
		//err = db.Where(sql, params...).First(&info).Error
		//return info, err
	}

	// adm 为空，仅根据 location 查
	sql := "city LIKE ? or district LIKE ?"
	params := []any{location + "%", location + "%"}
	err = geo.db.Where(sql, params...).First(&info).Error
	return info, err
}
