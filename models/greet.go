package models

type Greet struct {
	Model
	Sentence string `gorm:"uniqueIndex;type:varchar(255);not null;comment:句子" json:"sentence"`
	Author   string `gorm:"comment:作者" json:"author,omitempty"`
	Tags     string `gorm:"comment:标签" json:"tags,omitempty"`
}
