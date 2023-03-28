package model

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	UserID   uint `gorm:"index:idx_user"`
	User     User
	WhomId   uint `gorm:"index:idx_user2"`
	WhomUser User `gorm:"foreignKey:WhomId"`
}
