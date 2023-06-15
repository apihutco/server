package models

type GeoBank struct { // 国家
	Model
	// 国家
	Country string `gorm:"comment:国家;index" json:"country"`
	// 省份
	Province string `gorm:"comment:省份;index" json:"province"`
	// 城市
	City string `gorm:"comment:城市;index" json:"city"`
	// 区
	District string `gorm:"comment:区;index" json:"district"`
	// GeoID
	GeoID int64 `gorm:"comment:geoID;uniqueIndex" json:"geo_id"`
	// 时区
	TimeZone string `gorm:"comment:时区" json:"time_zone"`
	// UTC时区偏移
	UTCOffset string `gorm:"comment:UTC时区偏移" json:"utc_offset"`
	// 数据源
	//Sources []string `gorm:"comment:数据源" json:"sources"`
}
