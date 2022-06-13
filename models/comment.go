package models

import (
	"gorm.io/gorm"
	"gostars/utils/code"
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

func commentTableName() string {
	return "comments"
}

func CreateComment(data *Comment) int {
	err := db.Table(commentTableName()).Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func GetArticleComments(id int) ([]Comment, int) {
	var comments []Comment

	err := db.Table(commentTableName()).Where("article_id = ?", id).Where("status = ?", 1).Find(&comments).Error
	if err != nil {
		return comments, code.ERROR
	}

	return comments, code.SUCCESS
}

func GetArticleCommentCount(id int) int64 {
	var comment Comment
	var total int64

	err := db.Table(commentTableName()).Find(comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	if err != nil {
		return 0
	}

	return total
}

func DeleteComment(id uint) int {
	var comment Comment
	err := db.Table(commentTableName()).Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
