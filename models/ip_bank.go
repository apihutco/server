package models

import (
	"fmt"
	"time"
)

type IPBank struct {
	Model
	// IP
	IP string `gorm:"uniqueIndex;type:varchar(255);not null;comment:IP地址" json:"ip"`
	// 国家
	Country string `gorm:"comment:国家" json:"country"`
	// 省份
	Province string `gorm:"comment:省份" json:"province"`
	// 城市
	City string `gorm:"comment:城市" json:"city"`
	// 区
	District string `gorm:"comment:区" json:"district"`
	// 运营商
	ISP string `gorm:"comment:运营商" json:"isp"`
	// 地理位置
	Location string `gorm:"comment:地理位置" json:"location"`
	// 数据源
	Source string `gorm:"comment:数据源" json:"source"`
	// 缓存更新时间
	CacheTime time.Time `gorm:"-" json:"cache_time"`
}

func (i *IPBank) String() string {
	tmpl := "IP       : %s\nAddress  : %s %s %s\nISP      : %s\n"

	return fmt.Sprintf(tmpl,
		i.IP,
		i.Country,
		i.Province,
		i.City,
		i.ISP,
	)
}
