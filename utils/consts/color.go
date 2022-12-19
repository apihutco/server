package consts

type ColorCode string

const (
	ColorTransparent ColorCode = "transparent"
)

var colorStringMap = map[ColorCode]string{
	ColorTransparent: "transparent",
}

func (c ColorCode) String() string {
	return colorStringMap[c]
}
