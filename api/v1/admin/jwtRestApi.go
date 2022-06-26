package admin

import (
	"github.com/gin-gonic/gin"
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
		return
	}
	errCode = adminJwtService.DelRedisJwt(username)
	if errCode != code.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   errCode,
		"username": username,
		"message":  code.GetErrMsg(errCode),
	})
}
