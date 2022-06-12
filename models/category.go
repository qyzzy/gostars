package models

import (
	"fmt"
	"gorm.io/gorm"
	"gostars/utils/code"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20);not null" json:"name"`
}

func categoryTableName() string {
	return "categories"
}

func CreateCategory(data *Category) int {
	err := db.Table(categoryTableName()).Create(&data).Error
	fmt.Println(err)
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func CheckCategory(name string) int {
	var category Category
	db.Table(categoryTableName()).Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return code.ErrorCategoryNameUsed //2001
	}
	return code.SUCCESS
}
