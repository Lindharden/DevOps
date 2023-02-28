package model

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	WhoId  int
	WhomId int
}
