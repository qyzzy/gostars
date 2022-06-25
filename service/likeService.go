package service

import (
	"gostars/global"
	"gostars/models"
	"gostars/utils"
	"gostars/utils/code"
)

type LikeService struct {
}

func (likeService *LikeService) CreateLike(data *models.Like) int {
	err := global.GDb.Table(models.LikeTableName()).
		Create(&data).Error
	if err != nil {
		return code.ErrorLikeArticleFailed
	}
	return code.SUCCESS
}

func (likeService *LikeService) GetLikeUserIDList(articleID int) ([]int, int) {
	var likeUserIDList []int
	err := global.GDb.Table(models.LikeTableName()).Model(&models.Like{}).
		Where(map[string]interface{}{"article_id": articleID, "cancel": utils.IsLike}).
		Pluck("user_id", &likeUserIDList).Error
	if err != nil {
		return likeUserIDList, code.ERROR
	}
	return likeUserIDList, code.SUCCESS
}

func (likeService *LikeService) UpdateLike(articleID, userID, action int) int {
	err := global.GDb.Table(models.LikeTableName()).Model(&models.Like{}).
		Where(map[string]interface{}{"article_id": articleID, "user_id": userID}).
		Update("cancel", action).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (likeService *LikeService) GetLikeInfo(articleID, userID int) (models.Like, int) {
	var likeInfo models.Like
	err := global.GDb.Table(models.LikeTableName()).Model(&models.Like{}).
		Where(map[string]interface{}{"article_id": articleID, "user_id": userID}).
		First(&likeInfo).Error
	if err != nil {
		if err.Error() == "record not found" {
			return models.Like{}, code.ErrorDataNotFound
		}
		return likeInfo, code.ERROR
	}
	return likeInfo, code.SUCCESS
}

func (likeService *LikeService) GetLikeArticleList(userID int) ([]int, int) {
	var likeArticleList []int
	err := global.GDb.Table(models.LikeTableName()).Model(&models.Like{}).
		Where(map[string]interface{}{"user_id": userID, "cancel": utils.IsLike}).
		Pluck("article_id", likeArticleList).Error
	if err != nil {
		if err.Error() == "record not found" {
			return likeArticleList, code.ErrorDataNotFound
		}
		return likeArticleList, code.ERROR
	}
	return likeArticleList, code.SUCCESS
}
