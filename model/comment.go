package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
}
