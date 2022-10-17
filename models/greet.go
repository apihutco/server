package models

type Greet struct {
	Model
	Sentence string `gorm:"uniqueIndex;not null;comment:句子" json:"sentence"`
	Author   string `gorm:"comment:作者" json:"author"`
	Tags     string `gorm:"comment:标签" json:"tags"`
}
