package consts

import (
	"time"
)

type DayCode int

// Day：星期一 星期二 星期四 星期五 星期六 星期天

const (
	DayDefault DayCode = iota
	DayMonday
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)

var dayMap = map[DayCode]string{
	DayDefault:   DefaultCode.CN(),
	DayMonday:    "星期一",
	DayTuesday:   "星期二",
	DayWednesday: "星期三",
	DayThursday:  "星期四",
	DayFriday:    "星期五",
	DaySaturday:  "星期六",
	DaySunday:    "星期天",
}

// 换取星期几文字
func (d DayCode) String() string {
	return dayMap[d]
}

func GetDayCode() DayCode {
	switch time.Now().Weekday() {
	case time.Sunday:
		return DaySunday
	case time.Monday:
		return DayMonday
	case time.Tuesday:
		return DayTuesday
	case time.Wednesday:
		return DayWednesday
	case time.Thursday:
		return DayThursday
	case time.Friday:
		return DayFriday
	case time.Saturday:
		return DaySaturday
	default:
		return DayDefault
	}
}
