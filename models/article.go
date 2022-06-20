package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Summary      string `gorm:"type:varchar(200);not null" json:"summary"`
	Content      string `gorm:"type:longtext;not null" json:"content"`
	ClickCount   int    `gorm:"type:int;not null" json:"clickcount"`
	Status       int    `gorm:"type:tinyint;not null;default:0" json:"status"`
	AdminID      int    `gorm:"type:int;not null" json:"adminid"`
	IsOriginal   int    `gorm:"type:int;not null" json:"isoriginal"`
	Author       string `gorm:"type:varchar(100);not null" json:"author"`
	OpenComment  bool   `gorm:"type:tinyint;not null;default:0" json:"clickcount"`
	TagList      []Tag  `gorm:"type:text" json:"tagList"`
	Img          string `gorm:"type:varchar(100);not null" json:"img"`
	CommentCount int    `gorm:"type:int;not null" json:"commentcount"`
	CategoryID   int    `gorm:"type:int;not null" json:"categoryid"`
	CategoryName string `gorm:"type:varchar(50)" json:"categoryname"`
}

func ArticleTableName() string {
	return "articles"
}
