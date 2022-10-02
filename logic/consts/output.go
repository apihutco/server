package consts

type Output string

const (
	JSON Output = "json"
	PNG  Output = "png"
	JPG  Output = "jpg"
	JPEG Output = "jpeg"
	SVG  Output = "svg"
)

func (o Output) String() string {
	return string(o)
}
