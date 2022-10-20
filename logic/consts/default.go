package consts

type d struct{}

var (
	DefaultCode d = struct{}{}
)

func (i d) CN() string {
	return "通用"
}

func (i d) EN() string {
	return "default"
}
