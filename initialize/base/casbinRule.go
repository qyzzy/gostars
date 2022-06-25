package base

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"gostars/global"
)

var Casbin = new(casbin)

type casbin struct{}

var carbines = []gormadapter.CasbinRule{
	{Ptype: "p", V0: "888", V1: "/api/v1/users", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/v1/login", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/users", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/users", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/users", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/categories", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/articles", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/articles", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/articles", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/images", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/v1/admin/jwt/blacklist", V2: "POST"},
}

func (c *casbin) InitCasbinRule() error {
	_ = global.GDb.AutoMigrate(gormadapter.CasbinRule{})
	return global.GDb.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == int64(len(carbines)) {
			return nil
		}
		if err := tx.Create(carbines).Error; err != nil {
			return err
		}
		fmt.Println("init casbin rules success.")
		return nil
	})
}
