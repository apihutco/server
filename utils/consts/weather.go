package consts

type WeatherCode int

// Weather：晴 阴 云 雨 雾 雪 台风

var RepoWeather *weather

type weather struct {
	Default WeatherCode
	Sunny   WeatherCode
	Cloudy  WeatherCode
	Cloud   WeatherCode
	Wind    WeatherCode
	Rain    WeatherCode
	Fog     WeatherCode
	Typhoon WeatherCode
}

const (
	weatherDefault WeatherCode = iota
	weatherSunny
	weatherCloudy
	weatherCloud
	weatherWind
	weatherRain
	weatherFog
	weatherTyphoon
)

var weatherMap = map[WeatherCode]string{
	weatherDefault: DefaultCode.CN(),
	weatherSunny:   "晴",
	weatherCloudy:  "阴",
	weatherCloud:   "云",
	weatherWind:    "风",
	weatherRain:    "雨",
	weatherFog:     "雪",
	weatherTyphoon: "台风",
}

func init() {
	RepoWeather = &weather{
		Default: weatherDefault,
		Sunny:   weatherSunny,
		Cloudy:  weatherCloudy,
		Cloud:   weatherCloud,
		Wind:    weatherWind,
		Rain:    weatherRain,
		Fog:     weatherFog,
		Typhoon: weatherTyphoon,
	}
}

// 换取天气文字
func (w WeatherCode) String() string {
	return weatherMap[w]
}
