package consts

import (
	"time"
)

type SeasonCode int

// Season: 春 夏 秋 冬

const (
	SeasonDefault SeasonCode = iota
	SeasonSpring
	SeasonSummer
	SeasonAutumn
	SeasonWinter
)

var seasonMap = map[SeasonCode]string{
	SeasonDefault: "通用",
	SeasonSpring:  "春天",
	SeasonSummer:  "夏天",
	SeasonAutumn:  "秋天",
	SeasonWinter:  "冬天",
}

func (s SeasonCode) String() string {
	return seasonMap[s]
}

func GetSeasonCode() SeasonCode {
	m := time.Now().Month()
	switch {
	case m < 3 || m == 12:
		return SeasonWinter
	case m < 6:
		return SeasonSpring
	case m < 9:
		return SeasonSummer
	case m < 12:
		return SeasonAutumn
	default:
		return SeasonDefault
	}
}
