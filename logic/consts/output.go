package consts

import (
	"strings"
)

type OutputType string

const (
	JSON   OutputType = "json"
	PNG    OutputType = "png"
	JPG    OutputType = "jpg"
	JPEG   OutputType = "jpeg"
	SVG    OutputType = "svg"
	Base64 OutputType = "base64"
)

func (o OutputType) String() string {
	return string(o)
}

func CheckOutputType(in string, ot OutputType) bool {
	return strings.ToLower(in) == ot.String()
}
