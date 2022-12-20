package consts

import (
	"strings"
)

type OutputType string

var RepoOutput *output

type output struct {
	Default OutputType
	JSON    OutputType
	PNG     OutputType
	JPG     OutputType
	JPEG    OutputType
	SVG     OutputType
	Base64  OutputType
	Text    OutputType
}

var (
	outputDefault OutputType = OutputType(DefaultCode.EN())
	outputJSON    OutputType = "json"
	outputPNG     OutputType = "png"
	outputJPG     OutputType = "jpg"
	outputJPEG    OutputType = "jpeg"
	outputSVG     OutputType = "svg"
	outputBase64  OutputType = "base64"
	outputText    OutputType = "text"
)

func init() {
	RepoOutput = &output{
		Default: outputDefault,
		JSON:    outputJSON,
		PNG:     outputPNG,
		JPG:     outputJPG,
		JPEG:    outputJPEG,
		SVG:     outputSVG,
		Base64:  outputBase64,
		Text:    outputText,
	}
}

func (o OutputType) String() string {
	return string(o)
}

func (o *output) Is(str string, ot OutputType) bool {
	return strings.ToLower(str) == ot.String()
}

func (o *output) ToOutputType(outputStr string) OutputType {
	switch true {
	case o.Is(outputStr, outputJSON):
		return outputJSON
	case o.Is(outputStr, outputPNG):
		return outputPNG
	case o.Is(outputStr, outputJPG):
		return outputJPG
	case o.Is(outputStr, outputJPEG):
		return outputJPEG
	case o.Is(outputStr, outputSVG):
		return outputSVG
	case o.Is(outputStr, outputBase64):
		return outputBase64
	case o.Is(outputStr, outputText):
		return outputText
	default:
		return outputDefault
	}
}
