package config

type Open struct {
	*Tencent  `mapstructure:"tencent"`
	*Gaode    `mapstructure:"gaode"`
	*QWeather `mapstructure:"qweather"`
}

// Tencent 腾讯位置服务
type Tencent struct {
	Key string
}

// Gaode 高德开放平台
type Gaode struct {
	Key string
}

// QWeather 和风天气
type QWeather struct {
	Key string
}
