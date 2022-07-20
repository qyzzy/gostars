package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gostars/global"
	"gostars/initialize/base"
	"gostars/models"
	"gostars/utils"
	"time"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		fmt.Println("Connect database failed")
	}

	global.GDb = conn

	_ = global.GDb.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Tag{},
		&models.Article{},
		&models.Comment{},
		&models.JwtBlacklist{},
		&models.Image{},
		&models.Like{},
		&models.Level0NavMenu{},
		&models.Level1NavMenu{},
	)

	sqlDB, _ := global.GDb.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(10 * time.Second)

	_ = base.Casbin.InitCasbinRule()
	_ = base.Authority.InitAuthority()
	_ = base.NavMenu.InitLevel0NavMenu()
	_ = base.NavMenu.InitLevel1NavMenu()
}
