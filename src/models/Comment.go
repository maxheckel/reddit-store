package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Text string  `gorm:"size:255"`
	Parent *Comment
	Score int32
	Posted time.Time
	Children []*Comment
	Post Post
	Hash string   `gorm:"not null;unique;size:255"`
}
