package consts

// Platform：本地 腾讯 高德

type PlatformCode uint8

var RepoPlatform *platform

type platform struct {
	Local   PlatformCode
	Tencent PlatformCode
	Gaode   PlatformCode
}

const (
	platformLocal PlatformCode = iota + 1
	platformTencent
	platformGaode
)

var platformName = map[PlatformCode]string{
	platformLocal:   "local",
	platformTencent: "tencent",
	platformGaode:   "gaode",
}

func init() {
	RepoPlatform = &platform{
		Local:   platformLocal,
		Tencent: platformTencent,
		Gaode:   platformGaode,
	}
}

func (p PlatformCode) String() string {
	if pl, ok := platformName[p]; ok {
		return pl
	}
	return ""
}
