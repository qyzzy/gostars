package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Address string `gorm:"type:varchar(100)" json:"address"`
	UserID  int    `gorm:"type:int" json:"userid"`
	Status  int    `gorm:"default:1" json:"status"`
}

func ImageTableName() string {
	return "images"
}
