package consts

import (
	"time"
)

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

func GetTimeCode() TimeCode {
	hour := time.Now().Hour()
	switch {
	case hour < 4:
		return TimeLateNight
	case hour < 6:
		return TimeEarlyMorning
	case hour < 10:
		return TimeMorning
	case hour < 13:
		return TimeNoon
	case hour < 17:
		return TimeAfternoon
	case hour < 20:
		return TimeEvening
	case hour < 24:
		return TimeNight
	default:
		return TimeDefault
	}
}
