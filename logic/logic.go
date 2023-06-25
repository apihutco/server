package logic

import (
	"github.com/apihutco/server/logic/geo_bank"
	"github.com/apihutco/server/logic/weather"
)

func Init() {
	geo_bank.InitGeoCtrl()
	weather.Init()
}
