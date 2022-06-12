package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
}

func commentTableName() string {
	return "comments"
}

func CreateComment() {

}
