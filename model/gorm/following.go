package model

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	WhoId  uint
	WhomId uint
}
