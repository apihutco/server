package response

type Code uint64

const (
	CodeSuccess    Code = 2000
	CodeBaeRequest Code = 4000
	CodeError      Code = 5000 + iota

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
