package models

type Weather struct {
	Model
	// 地区
	Location string `gorm:"comment:地区" json:"location"`
	// 数据观测时间
	ObsTime string `gorm:"comment:数据观测时间" json:"obs_time,omitempty"`
	FxDate  string `json:"fx_date,omitempty"`
	// 温度
	Temp string `gorm:"comment:温度" json:"temp,omitempty"`
	// 最低温度
	MinTemp string `json:"min_temp,omitempty"`
	// 最高温度
	MaxTemp string `json:"max_temp,omitempty"`
	// 体感温度
	FeelsLike string `gorm:"comment:体感温度" json:"feels_like,omitempty"`
	// 描述
	Text      string `gorm:"comment:描述" json:"text,omitempty"`
	TextDay   string `gorm:"comment:描述" json:"text_day,omitempty"`
	TextNight string `gorm:"comment:描述" json:"text_night,omitempty"`
	// 360度风向
	Wind360 string `gorm:"comment:360度风向" json:"wind360"`
	// 风向
	WindDir string `gorm:"comment:风向" json:"wind_dir"`
	// 相对湿度
	Humidity string `gorm:"comment:相对湿度" json:"humidity"`
	// 当前小时累计降水量，毫米
	Precip string `gorm:"comment:当前小时累计降水量" json:"precip"`
	// 可视化界面
	FxLink string `gorm:"comment:可视化界面" json:"fx_link,omitempty"`
	// 数据源
	Sources []string `gorm:"comment:数据源" json:"sources,omitempty"`
}
