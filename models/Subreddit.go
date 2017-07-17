package models

import "github.com/jinzhu/gorm"

type Subreddit struct {
	gorm.Model
	Name string
	Subreddit string
	Posts []Post
}

