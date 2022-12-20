package bleve

import (
	"errors"
)

var (
	ErrorIndexEmpty = errors.New("索引为空")
	ErrorNotFound   = errors.New("无匹配数据")
)
