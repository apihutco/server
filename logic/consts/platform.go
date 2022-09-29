package consts

type PlatformID uint8

const (
	Local PlatformID = iota + 1
	Tencent
	Gaode
)

var PlatformName = map[PlatformID]string{
	Local:   "local",
	Tencent: "tencent",
	Gaode:   "gaode",
}

func (p PlatformID) Name() string {
	if pl, ok := PlatformName[p]; ok {
		return pl
	}
	return ""
}
