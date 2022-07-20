package models

import "gorm.io/gorm"

type NavMenu struct {
	gorm.Model
	Name           string `json:"name"`
	CategoryLevel0 int    `json:"categoryLevel0"`
	CategoryLevel1 int    `json:"categoryLevel1"`
	Path           string `json:"path"`
}
