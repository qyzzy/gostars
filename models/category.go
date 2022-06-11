package models

import (
	"gorm.io/gorm"
	"gostars/utils/code"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20);not null" json:"name"`
}

func categoryTableName() string {
	return "categorys"
}

func CreateCategory(data *Category) int {
	err := db.Table(categoryTableName()).Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
