package logic

import (
	"apihut-server/logic/geo_bank"
	"apihut-server/logic/weather"
)

func Init() {
	geo_bank.InitGeoCtrl()
	weather.Init()
}
