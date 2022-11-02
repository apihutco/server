package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code uint64

const (
	CodeSuccess Code = 2000
	CodeError   Code = 5000 + iota

	ErrorFormat
	ErrorBind

	// IP定位
	ErrorIPUnableToLocate
	// 一句招呼

	// 哈希头像
	ErrorAvatarGenerate

	// 协议测试
	ErrorProtocolReadBody
	ErrorProtocolUnmarshal
	ErrorProtocolWsUpgrade
)

var CodeMsgMap = map[Code]string{
	CodeSuccess:            "成功",
	CodeError:              "失败",
	ErrorFormat:            "格式错误",
	ErrorBind:              "参数绑定失败",
	ErrorIPUnableToLocate:  "IP无法定位",
	ErrorAvatarGenerate:    "头像生成失败",
	ErrorProtocolReadBody:  "读取Body失败",
	ErrorProtocolUnmarshal: "序列化失败",
	ErrorProtocolWsUpgrade: "协议升级失败",
}

func (c Code) Msg() string {
	if msg, ok := CodeMsgMap[c]; ok {
		return msg
	}
	return ""
}

func Success(c *gin.Context) {
	SuccessWithData(c, nil)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	code := CodeSuccess
	JSON(c, code, code.Msg(), data)
}

func Error(c *gin.Context) {
	ErrorWithMsg(c, CodeError.Msg())
}

func ErrorWithCode(c *gin.Context, code Code) {
	JSON(c, code, code.Msg(), nil)
}

func ErrorWithMsg(c *gin.Context, msg string) {
	code := CodeError
	JSON(c, code, msg, nil)
}

func ErrorWithData(c *gin.Context, data interface{}) {
	code := CodeError
	JSON(c, code, code.Msg(), data)
}

func ErrorWithMsgAndData(c *gin.Context, msg string, data interface{}) {
	code := CodeError
	JSON(c, code, msg, data)
}

func JSON(c *gin.Context, code Code, msg string, data interface{}) {
	c.JSON(http.StatusOK, &Body{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
