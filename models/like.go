package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	ArticleID int  `json:"articleid"`
	UserID    int  `json:"userid"`
	Cancel    int8 `json:"cancel"`
}

func LikeTableName() string {
	return "likes"
}
