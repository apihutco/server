package consts

type WeatherCode int

// Weather：晴 阴 云 雨 雾 雪 台风

const (
	WeatherDefault WeatherCode = iota
	WeatherSunny
	WeatherCloudy
	WeatherCloud
	WeatherWind
	WeatherRain
	WeatherFog
	WeatherTyphoon
)

var weatherMap = map[WeatherCode]string{
	WeatherDefault: DefaultCode.CN(),
	WeatherSunny:   "晴",
	WeatherCloudy:  "阴",
	WeatherCloud:   "云",
	WeatherWind:    "风",
	WeatherRain:    "雨",
	WeatherFog:     "雪",
	WeatherTyphoon: "台风",
}

// 换取天气文字
func (w WeatherCode) String() string {
	return weatherMap[w]
}
