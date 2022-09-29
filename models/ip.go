package models

import (
	"gorm.io/gorm"
)

type IPBank struct {
	gorm.Model
	// IP
	IP string
	// 国家
	Country string
	// 省份
	Province string
	// 城市
	City string
	// 区
	District string
	// 运营商
	ISP string
	// 地理位置
	Location string
	// 数据源
	Source string
}
