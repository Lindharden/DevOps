package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Email      string
	PwHash     string
	Messages   []Message
	Followings []Following `gorm:"foreignKey:WhoId"`
}
