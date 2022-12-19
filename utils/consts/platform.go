package consts

type PlatformCode uint8

// Platform：本地 腾讯 高德

const (
	Local PlatformCode = iota + 1
	Tencent
	Gaode
)

var PlatformName = map[PlatformCode]string{
	Local:   "local",
	Tencent: "tencent",
	Gaode:   "gaode",
}

func (p PlatformCode) String() string {
	if pl, ok := PlatformName[p]; ok {
		return pl
	}
	return ""
}
