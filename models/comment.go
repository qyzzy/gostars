package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    int    `gorm:"type:int;not null" json:"userid"`
	ArticleID int    `gorm:"type:int;not null" json:"articleid"`
	Title     string `gorm:"type:varchar(50);not null" json:"title"`
	Content   string `gorm:"type:varchar(500);not null" json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2" json:"status"`
	IPSource  string `gorm:"type:varchar(100);not null" json:"ipsource"`
}

func CommentTableName() string {
	return "comments"
}
