package consts

import (
	"time"
)

// Time：通用 凌晨 早晨 中午 下午 傍晚 晚上  深夜

type TimeCode int

var RepoTime *_time

type _time struct {
	Default      TimeCode
	EarlyMorning TimeCode
	Morning      TimeCode
	Noon         TimeCode
	Afternoon    TimeCode
	Evening      TimeCode
	Night        TimeCode
	LateNight    TimeCode
}

const (
	timeDefault TimeCode = iota
	timeEarlyMorning
	timeMorning
	timeNoon
	timeAfternoon
	timeEvening
	timeNight
	timeLateNight
)

var timeMap = map[TimeCode]string{
	timeDefault:      DefaultCode.CN(),
	timeEarlyMorning: "凌晨",
	timeMorning:      "早上",
	timeNoon:         "中午",
	timeAfternoon:    "下午",
	timeEvening:      "傍晚",
	timeNight:        "晚上",
	timeLateNight:    "深夜",
}

func init() {
	RepoTime = &_time{
		Default:      timeDefault,
		EarlyMorning: timeEarlyMorning,
		Morning:      timeMorning,
		Noon:         timeNoon,
		Afternoon:    timeAfternoon,
		Evening:      timeEvening,
		Night:        timeNight,
		LateNight:    timeLateNight,
	}
}

// 换取时间段文字
func (c TimeCode) String() string {
	return timeMap[c]
}

func (*_time) Now() TimeCode {
	hour := time.Now().Hour()
	switch {
	case hour < 4:
		return timeLateNight
	case hour < 6:
		return timeEarlyMorning
	case hour < 10:
		return timeMorning
	case hour < 13:
		return timeNoon
	case hour < 17:
		return timeAfternoon
	case hour < 20:
		return timeEvening
	case hour < 24:
		return timeNight
	default:
		return timeDefault
	}
}
