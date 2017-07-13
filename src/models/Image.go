package models

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Path string  `gorm:"size:255"`
	OriginalUrl string  `gorm:"size:255"`
	Post Post
}

