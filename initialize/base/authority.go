package base

import (
	"fmt"
	"gorm.io/gorm"
	"gostars/global"
	"gostars/models"
	"time"
)

var Authority = new(authority)

type authority struct{}

var authorities = []models.Authority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityID: "888", AuthorityName: "admin"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityID: "777", AuthorityName: "user"},
}

func (a *authority) InitAuthority() error {
	_ = global.GDb.AutoMigrate(models.Authority{})
	return global.GDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(authorities).Error; err != nil {
			return err
		}
		fmt.Println("init casbin rules success.")
		return nil
	})
}
