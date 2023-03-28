package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  uint
	User    User
	Text    string
	PubDate int64 `gorm:"index:,sort:desc"`
	Flagged int
}
