package base

import (
	"gorm.io/gorm"
	"gostars/global"
	"gostars/models"
)

var NavMenu = new(navMenu)

type navMenu struct{}

var navMenus = []models.NavMenu{
	{Name: "Home", CategoryLevel0: 0, CategoryLevel1: 0, Path: "/api/v1/articles"},
	{Name: "Blog", CategoryLevel0: 0, CategoryLevel1: 0, Path: "/api/v1/articles"},
	{Name: "About", CategoryLevel0: 0, CategoryLevel1: 0, Path: "/api/v1/about"},
	{Name: "NetHistory", CategoryLevel0: 103, CategoryLevel1: 0, Path: "/api/v1/history"},
	{Name: "RSS", CategoryLevel0: 103, CategoryLevel1: 0, Path: "/api/v1/rss"},
}

func (a *navMenu) InitNavMenu() error {
	_ = global.GDb.AutoMigrate(models.NavMenu{})
	return global.GDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(navMenus).Error; err != nil {
			return err
		}
		return nil
	})
}
