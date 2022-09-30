package models

// Gravatar: s,d,f,r

type AvatarReq struct {
	Hash      string
	Size      int    `form:"size"`
	S         int    `form:"s"` // Gravatar
	UDefault  string `form:"default"`
	D         string `form:"d"`         // Gravatar
	Random    bool   `from:"random"`    // 随机
	Density   int    `form:"density"`   // 密度
	Namespace string `form:"namespace"` // 命名空间
	N         string `form:"namespace"` // 命名空间
}

func NewAvatar(hash string) AvatarReq {
	return AvatarReq{
		Hash:      hash,
		Size:      0,
		Random:    false,
		Density:   1,
		Namespace: "",
	}
}

func (a *AvatarReq) GetSize() int {
	if a.Size != 0 {
		return a.Size
	}
	if a.S != 0 {
		return a.S
	}
	return 4
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
	return a.Density
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
