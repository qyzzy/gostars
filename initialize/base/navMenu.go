package base

import (
	"gorm.io/gorm"
	"gostars/global"
	"gostars/models"
)

var NavMenu = new(navMenu)

type navMenu struct{}

var level0NavMenus = []models.Level0NavMenu{
	{Name: "Home", Path: "/api/v1/home", IsFather: false},
	{Name: "Blog", Path: "/api/v1/articles", IsFather: false},
	{Name: "About", Path: "/api/v1/about", IsFather: true},
}

var level1NavMenus = []models.Level1NavMenu{
	{Name: "NetHistory", Path: "/api/v1/net/history", FatherID: 3, IsFather: false},
	{Name: "RSS", Path: "/api/v1/rss", FatherID: 3, IsFather: false},
}

func (a *navMenu) InitLevel0NavMenu() error {
	_ = global.GDb.AutoMigrate(models.Level0NavMenu{})
	return global.GDb.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&models.Level0NavMenu{}).RowsAffected == int64(len(level0NavMenus)) {
			return nil
		}
		if err := tx.Create(level0NavMenus).Error; err != nil {
			return err
		}
		return nil
	})
}

func (a *navMenu) InitLevel1NavMenu() error {
	_ = global.GDb.AutoMigrate(models.Level1NavMenu{})
	return global.GDb.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&models.Level1NavMenu{}).RowsAffected == int64(len(level1NavMenus)) {
			return nil
		}
		if err := tx.Create(level1NavMenus).Error; err != nil {
			return err
		}
		return nil
	})
}
