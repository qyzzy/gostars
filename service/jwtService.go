package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gostars/global"
	"gostars/models"
	"gostars/utils"
	"gostars/utils/code"
	"time"
)

type JwtService struct {
}

/**
	MySQL Part
**/
func (jwtService *JwtService) AddJwtTokenToBlacklist(jwtList *models.JwtBlacklist) int {
	err := global.GDb.Table(models.JwtBlacklistTableName()).
		Create(&jwtList).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.GDb.Table(models.JwtBlacklistTableName()).
		Where("jwt = ?", jwt).First(&models.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

/**
	Redis Part
**/
func (jwtService *JwtService) GetRedisJwt(username string) (errCode int, redisJwt string) {
	redisJwt, err := global.GRedisGroup[0].Get(context.Background(), username).Result()
	if err != nil {
		return code.ErrorRedisGetFailed, redisJwt
	}
	return code.SUCCESS, redisJwt
}

func (jwtService *JwtService) SetRedisJwt(jwt, username string) (errCode int) {
	timer := time.Duration(utils.JwtExpireTime) * time.Second
	err := global.GRedisGroup[0].Set(context.Background(), username, jwt, timer).Err()
	if err != nil {
		return code.ErrorRedisSaveFailed
	}
	return code.SUCCESS
}

func (jwtService *JwtService) DelRedisJwt(username string) (errCode int) {
	err := global.GRedisGroup[0].Del(context.Background(), username).Err()
	if err != nil {
		return code.ErrorRedisDeleteFailed
	}
	return code.SUCCESS
}
