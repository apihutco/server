package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	statusCode int
	body       *Body
	c          *gin.Context
}

func NewResponse(c *gin.Context, statusCode int, code Code) *Response {
	body := new(Body)
	body.Code = code
	return &Response{
		statusCode: statusCode,
		body:       body,
		c:          c,
	}
}

func Success(c *gin.Context) *Response {
	return NewResponse(c, http.StatusOK, CodeSuccess)
}

func Error(c *gin.Context) *Response {
	return NewResponse(c, http.StatusInternalServerError, CodeError)
}

func BadRequest(c *gin.Context) *Response {
	return NewResponse(c, http.StatusBadRequest, CodeBaeRequest)
}

func (r *Response) Code(code Code) *Response {
	r.body.Code = code
	return r
}

func (r *Response) Msg(msg string) *Response {
	r.body.Msg = msg
	return r
}

func (r *Response) Data(data any) *Response {
	r.body.Data = data
	return r
}

func (r *Response) checkAndFix() {
	if r.body.Code == 0 {
		r.body.Code = CodeSuccess
	}
	if len(r.body.Msg) == 0 {
		r.body.Msg = r.body.Code.Msg()
	}
}

// JSON return code,msg and data
func (r *Response) JSON() {
	r.checkAndFix()
	r.c.JSON(r.statusCode, r.body)
}

// Pure return data with code and msg
func (r *Response) Pure() {
	r.checkAndFix()
	r.c.JSON(r.statusCode, r.body.Data)
}

// String return as text
func (r *Response) String() {
	r.checkAndFix()
	str, ok := r.body.Data.(string)
	if !ok {
		str = ErrorProtocolUnmarshal.Msg()
	}
	r.c.String(r.statusCode, str)
}
