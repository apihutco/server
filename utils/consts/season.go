package consts

import (
	"time"
)

// Season: 春 夏 秋 冬

type SeasonCode int

var RepoSeason *season

type season struct {
	Default SeasonCode
	Spring  SeasonCode
	Summer  SeasonCode
	Autumn  SeasonCode
	Winter  SeasonCode
}

const (
	seasonDefault SeasonCode = iota
	seasonSpring
	seasonSummer
	seasonAutumn
	seasonWinter
)

var seasonMap = map[SeasonCode]string{
	seasonDefault: DefaultCode.CN(),
	seasonSpring:  "春天",
	seasonSummer:  "夏天",
	seasonAutumn:  "秋天",
	seasonWinter:  "冬天",
}

func init() {
	RepoSeason = &season{
		Default: seasonDefault,
		Spring:  seasonSpring,
		Summer:  seasonSummer,
		Autumn:  seasonAutumn,
		Winter:  seasonWinter,
	}
}

func (s SeasonCode) String() string {
	return seasonMap[s]
}

func (*season) Today() SeasonCode {
	m := time.Now().Month()
	switch {
	case m < 3 || m == 12:
		return seasonWinter
	case m < 6:
		return seasonSpring
	case m < 9:
		return seasonSummer
	case m < 12:
		return seasonAutumn
	default:
		return seasonDefault
	}
}
