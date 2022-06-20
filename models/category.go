package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20);not null" json:"name"`
}

func CategoryTableName() string {
	return "categories"
}
