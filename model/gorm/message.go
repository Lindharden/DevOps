package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID  int
	User    User
	Text    string
	PubDate int64
	Flagged int
}
