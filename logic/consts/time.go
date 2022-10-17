package consts

type TimeCode int

// Time：通用 凌晨 早晨 中午 下午 傍晚 晚上  深夜

const (
	TimeDefault TimeCode = iota
	TimeEarlyMorning
	TimeMorning
	TimeNoon
	TimeAfternoon
	TimeEvening
	TimeNight
	TimeLateNight
)

var timeMap = map[TimeCode]string{
	TimeDefault:      "通用",
	TimeEarlyMorning: "凌晨",
	TimeMorning:      "早上",
	TimeNoon:         "中午",
	TimeAfternoon:    "下午",
	TimeEvening:      "傍晚",
	TimeNight:        "晚上",
	TimeLateNight:    "深夜",
}

// 换取时间段文字
func (c TimeCode) String() string {
	return timeMap[c]
}
