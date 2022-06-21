package models

import (
	"gorm.io/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt" json:"text"`
}

func JwtBlacklistTableName() string {
	return "jwtblacklists"
}
