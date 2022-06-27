package service

import (
	"context"
	"gostars/global"
	"gostars/models"
	"gostars/utils"
	"gostars/utils/code"
	"time"
)

type LikeService struct {
}

/**
	MySQL Part
**/
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

/**
	Redis Part
**/
func (likeService *LikeService) ExistsRedisLikeUserID(strUserID string) int {
	res, err := global.GRedisGroup[1].Exists(context.Background(), strUserID).Result()
	if err != nil {
		return code.ERROR
	}
	if res == 0 {
		return code.ErrorRedisKeyNotExist
	}
	return code.SUCCESS
}

func (likeService *LikeService) SAddRedisLikeUserID(strUserID string, articleID int) int {
	_, err := global.GRedisGroup[1].SAdd(context.Background(), strUserID, articleID).Result()
	if err != nil {
		return code.ErrorRedisSaveFailed
	}
	return code.SUCCESS
}

func (likeService *LikeService) DelRedisLikeUserID(strUserID string) int {
	err := global.GRedisGroup[1].Del(context.Background(), strUserID).Err()
	if err != nil {
		return code.ErrorRedisDeleteFailed
	}
	return code.SUCCESS
}

func (likeService *LikeService) ExpireRedisLikeUserID(strUserID string) int {
	err := global.GRedisGroup[1].Expire(
		context.Background(),
		strUserID,
		time.Duration(utils.Week)*time.Second,
	)
	if err != nil {
		return code.ErrorRedisExpireFailed
	}
	return code.SUCCESS
}
