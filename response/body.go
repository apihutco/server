package response

type Body struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
