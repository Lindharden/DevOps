package model

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	UserID   uint
	User     User
	WhomId   uint
	WhomUser User `gorm:"foreignKey:WhomId"`
}
