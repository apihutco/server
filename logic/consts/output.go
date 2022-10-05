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
	Text   OutputType = "text"
)

func (o OutputType) String() string {
	return string(o)
}

func IsOutputType(in string, ot OutputType) bool {
	return strings.ToLower(in) == ot.String()
}

func CaseOutputType(in string, ot OutputType) string {
	if IsOutputType(in, ot) {
		return in
	}
	return ""
}

func ToOutputType(in string) OutputType {
	return OutputType(in)
}
