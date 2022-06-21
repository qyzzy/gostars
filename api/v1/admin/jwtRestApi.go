package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
)

func (jwtApi *JwtApi) AddBlackList(c *gin.Context) {
	var username string
	var jwtList = new(models.JwtBlacklist)
	var token string
	var errCode int
	username = c.Query("username")

	errCode, token = adminJwtService.GetRedisJwt(username)
	if errCode != code.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		})
		return
	}

	jwtList.Jwt = token
	errCode = adminJwtService.AddJwtTokenToBlacklist(jwtList)
	if errCode != code.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		})
	}
	err := global.GRedis.Del(context.Background(), username).Err()
	if err != nil {
		errCode = code.ErrorRedisDeleteFailed
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   errCode,
		"username": username,
		"message":  code.GetErrMsg(errCode),
	})
}
