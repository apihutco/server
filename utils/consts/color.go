package consts

import "strings"

type ColorCode string

var RepoColor *color

type color struct {
	Transparent ColorCode
}

const (
	colorTransparent ColorCode = "transparent"
)

var colorStringMap = map[ColorCode]string{
	colorTransparent: "transparent",
}

func init() {
	RepoColor = &color{
		Transparent: colorTransparent,
	}
}

func (c ColorCode) String() string {
	return colorStringMap[c]
}

func (c *color) Is(str string, colorCode ColorCode) bool {
	return strings.ToLower(str) == colorCode.String()
}
