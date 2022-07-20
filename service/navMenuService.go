package service

import (
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type Level0NavMenuService struct {
}

func (level0NavMenuService *Level0NavMenuService) GetLevel0NavMenus() (*[]models.Level0NavMenu, int) {
	var Level0NavMenus *[]models.Level0NavMenu
	err := global.GDb.Table(models.Level0NavMenuTableName()).
		Select("*").Find(&Level0NavMenus).Error
	if err != nil {
		return nil, code.ERROR
	}
	return Level0NavMenus, code.SUCCESS
}

type Level1NavMenuService struct {
}

func (level1NavMenuService *Level1NavMenuService) GetLevel1NavMenus() (*[]models.Level1NavMenu, int) {
	var Level1NavMenus *[]models.Level1NavMenu
	err := global.GDb.Table(models.Level1NavMenuTableName()).
		Select("*").Find(&Level1NavMenus).Error
	if err != nil {
		return nil, code.ERROR
	}
	return Level1NavMenus, code.SUCCESS
}
