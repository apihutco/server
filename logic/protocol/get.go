package protocol

import (
	"fmt"
	"net/http"
	"strings"
)

const formatString = "%s:%s "

// 单个参数以string返回，多个重复key的参数以数组返回

func ParamsToJSON(req *http.Request) map[string]any {
	res := make(map[string]any)

	for key, values := range req.URL.Query() {
		// 遍历重复key下的values
		for _, value := range values {
			// 检查key是否存在
			if _, ok := res[key]; ok {
				// 根据类型操作
				switch res[key].(type) {
				case string:
					tmp := res[key]
					delete(res, key)
					res[key] = make([]string, 0)
					res[key] = append((res[key]).([]string), tmp.(string))
					res[key] = append((res[key]).([]string), value)
				case []string:
					res[key] = append((res[key]).([]string), value)
				}
			} else {
				// key不存在则初始化
				res[key] = value
			}
		}
	}

	return res
}

func ParamsToString(req *http.Request) string {
	builder := strings.Builder{}

	for key, values := range req.URL.Query() {
		for _, value := range values {
			builder.WriteString(fmt.Sprintf(formatString, key, value))
		}
	}

	return builder.String()
}
