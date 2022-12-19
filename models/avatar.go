package models

import (
	"apihut-server/utils/consts"
)

// Gravatar: s,d,f,r

type AvatarReq struct {
	Hash      string            // 哈希（不受Query参数值控制）
	Block     int               `form:"block"`   // 块数量
	Size      int               `form:"size"`    // Gravatar,图像大小
	S         int               `form:"s"`       // Gravatar,图像大小-缩写
	UDefault  string            `form:"default"` // Gravatar,默认图片
	D         string            `form:"d"`       // Gravatar,默认图片-缩写
	Random    bool              // 随机(不受参数值控制，只要存在random参数，就视为随机模式)
	Density   int               `form:"density"`   // 密度
	Namespace string            `form:"namespace"` // 命名空间
	N         string            `form:"namespace"` // 命名空间-缩写
	Output    consts.OutputType `form:"output"`    // 输出格式
	O         consts.OutputType `form:"O"`         // 输出格式-简写
	Quality   int               `form:"quality"`   // 图像质量（仅jpg，jpeg可用）
	// Pixels    int           `form:"pixels"`    // 图像像素（图片大小）
	BackgroundColor string `form:"backgroundcolor"` // 背景颜色
	BGColor         string `form:"bgcolor"`         // 背景颜色-缩写
	FillColor       string `form:"fillcolor"`       // 填充颜色
	FColor          string `form:"fcolor"`          // 填充颜色-缩写
}

func NewAvatar(hash string) AvatarReq {
	return AvatarReq{
		Hash:      hash,
		Size:      0,
		Density:   1,
		Namespace: "",
	}
}

func (a *AvatarReq) GetSize() int {
	if a.Size > 0 {
		return a.Size
	}
	if a.S > 0 {
		return a.S
	}
	return 32
}

func (a *AvatarReq) GetDefault() string {
	if len(a.UDefault) != 0 {
		return a.UDefault
	}
	return a.D
}

func (a *AvatarReq) SetHash(h string) {
	a.Hash = h
}

func (a *AvatarReq) GetHash() string {
	return a.Hash
}

func (a *AvatarReq) GetDensity() int {
	if a.Density > 1 {
		return a.Density
	}
	return 1
}

func (a *AvatarReq) GetNamespace() string {
	if len(a.Namespace) != 0 {
		return a.Namespace
	}
	if len(a.N) != 0 {
		return a.N
	}
	return "apihut"
}

func (a *AvatarReq) GetOutput() consts.OutputType {
	if len(a.Output) != 0 {
		return a.Output
	}
	if len(a.O) != 0 {
		return a.O
	}
	return "png"
}

func (a *AvatarReq) GetQuality() int {
	if a.Quality != 0 {
		return a.Quality
	}
	return 100
}

// func (a *AvatarReq) GetPixels() int {
// 	if a.Pixels != 0 {
// 		return a.Pixels
// 	}
// 	return 10
// }

func (a *AvatarReq) GetBlock() int {
	if a.Block > 4 {
		return a.Block
	}
	return 4
}

func (a *AvatarReq) GetBackgroundColor() string {
	if len(a.BackgroundColor) > 0 {
		return a.BackgroundColor
	}
	return a.BGColor
}

func (a *AvatarReq) GetFillColor() string {
	if len(a.FillColor) > 0 {
		return a.FillColor
	}
	return a.FColor
}
