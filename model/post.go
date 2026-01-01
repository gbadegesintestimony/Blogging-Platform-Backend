package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string
	Content    string
	UserID     uint
	CategoryID uint
	comments   []Comment
}
