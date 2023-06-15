package mysql

import (
	"apihut-server/models"
)

func CreateGeoBank(geoInfo *models.GeoBank) error {
	db.Where("")
	return db.Create(&geoInfo).Error
}

func GetGeoBank(location string, adm string) (*models.GeoBank, error) {
	info := new(models.GeoBank)
	var err error

	if adm != "" {
		sql := "(district LIKE ? or city LIKE ?) and (city LIKE ? or province LIKE ?)"
		params := []any{location + "%", location + "%", adm + "%", adm + "%"}
		err = db.Where(sql, params...).First(&info).Error
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
	err = db.Where(sql, params...).First(&info).Error
	return info, err
}
