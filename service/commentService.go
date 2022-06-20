package service

import (
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type CommentService struct {
}

func (commentService *CommentService) CreateComment(data *models.Comment) int {
	err := global.GDb.Table(models.CommentTableName()).
		Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (commentService *CommentService) GetArticleComments(id int) ([]models.Comment, int) {
	var comments []models.Comment

	err := global.GDb.Table(models.CommentTableName()).
		Where("article_id = ?", id).Where("status = ?", 1).Find(&comments).Error
	if err != nil {
		return comments, code.ERROR
	}

	return comments, code.SUCCESS
}

func (commentService *CommentService) GetArticleCommentCount(id int) int64 {
	var comment models.Comment
	var total int64

	err := global.GDb.Table(models.CommentTableName()).
		Find(comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	if err != nil {
		return 0
	}

	return total
}

func (commentService *CommentService) DeleteComment(id uint) int {
	var comment models.Comment
	err := global.GDb.Table(models.CommentTableName()).
		Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}
