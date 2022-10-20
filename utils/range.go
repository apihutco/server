package utils

import (
	"math/rand"
	"time"
)

func GetRange(i int) int {
	// 赋予随机种子
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}
