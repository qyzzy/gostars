package models

import "gorm.io/gorm"

type Level0NavMenu struct {
	gorm.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsFather bool   `json:"isFather"`
}

type Level1NavMenu struct {
	gorm.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	FatherID int    `json:"fatherID"`
	IsFather bool   `json:"isFather"`
}

func Level0NavMenuTableName() string {
	return "level0_nav_menus"
}

func Level1NavMenuTableName() string {
	return "level1_nav_menus"
}
