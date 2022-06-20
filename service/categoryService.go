package service

import (
	"fmt"
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type CategoryService struct {
}

func (categoryService *CategoryService) CreateCategory(data *models.Category) int {
	err := global.GDb.Table(models.CategoryTableName()).
		Create(&data).Error
	fmt.Println(err)
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (categoryService *CategoryService) CheckCategory(name string) int {
	var category models.Category
	global.GDb.Table(models.CategoryTableName()).
		Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return code.ErrorCategoryNameUsed //2001
	}
	return code.SUCCESS
}
