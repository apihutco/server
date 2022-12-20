package consts

import (
	"time"
)

// Day：星期一 星期二 星期四 星期五 星期六 星期天

type DayCode int

var RepoDay *day

type day struct {
	Default   DayCode
	Monday    DayCode
	Tuesday   DayCode
	Wednesday DayCode
	Thursday  DayCode
	Friday    DayCode
	Saturday  DayCode
	Sunday    DayCode
}

const (
	dayDefault DayCode = iota
	dayMonday
	dayTuesday
	dayWednesday
	dayThursday
	dayFriday
	daySaturday
	daySunday
)

var dayMap = map[DayCode]string{
	dayDefault:   DefaultCode.CN(),
	dayMonday:    "星期一",
	dayTuesday:   "星期二",
	dayWednesday: "星期三",
	dayThursday:  "星期四",
	dayFriday:    "星期五",
	daySaturday:  "星期六",
	daySunday:    "星期天",
}

func init() {
	RepoDay = &day{
		Default:   dayDefault,
		Monday:    dayMonday,
		Tuesday:   dayTuesday,
		Wednesday: dayWednesday,
		Thursday:  dayThursday,
		Friday:    dayFriday,
		Saturday:  daySaturday,
		Sunday:    daySunday,
	}
}

// 换取星期几文字
func (d DayCode) String() string {
	return dayMap[d]
}

func (*day) Today() DayCode {
	switch time.Now().Weekday() {
	case time.Sunday:
		return daySunday
	case time.Monday:
		return dayMonday
	case time.Tuesday:
		return dayTuesday
	case time.Wednesday:
		return dayWednesday
	case time.Thursday:
		return dayThursday
	case time.Friday:
		return dayFriday
	case time.Saturday:
		return daySaturday
	default:
		return dayDefault
	}
}
