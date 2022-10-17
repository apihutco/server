package consts

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
