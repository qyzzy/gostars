package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	jwttoken "gostars/utils/jwt"
	"net/http"
	"time"
)

func (userApi *UserApi) LoginFront(c *gin.Context) {
	var formData models.User
	var errCode int
	_ = c.ShouldBindJSON(&formData)

	formData, errCode = webUserService.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": code.GetErrMsg(errCode),
	})
}

func (userApi *UserApi) Login(c *gin.Context) {
	var formData models.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var errCode int
	var loginErrCode int

	formData, loginErrCode = webUserService.CheckLogin(formData.Username, formData.Password)

	errCode, token = webJwtService.GetRedisJwt(formData.Username)
	// judge jwt token in blacklist
	if !webJwtService.IsBlacklist(token) {
		c.JSON(http.StatusOK, gin.H{
			"status":  code.ErrorTokenInBlacklist,
			"message": code.GetErrMsg(code.ErrorTokenInBlacklist),
		})
		return
	}

	if errCode == code.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code.SUCCESS,
			"data":    formData.Username,
			"id":      formData.ID,
			"role":    formData.Role,
			"message": code.GetErrMsg(code.SUCCESS),
			"token":   token,
		})
		return
	}

	if loginErrCode == code.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  loginErrCode,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": code.GetErrMsg(loginErrCode),
			"token":   token,
		})
	}

}

func setToken(c *gin.Context, user models.User) {
	j := jwttoken.NewJwt()
	claims := jwttoken.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 6000,
			Issuer:    "GoStars",
		},
	}

	token, err := j.CreateToken(claims)

	// save jwt token to redis
	errCode := webJwtService.SetRedisJwt(token, user.Username)
	if errCode != code.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code.ErrorRedisSaveFailed,
			"message": code.GetErrMsg(code.ErrorRedisSaveFailed),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code.ErrorTokenCreateFailed,
			"message": code.GetErrMsg(code.ERROR),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code.SUCCESS,
		"data":    user.Username,
		"id":      user.ID,
		"message": code.GetErrMsg(200),
		"token":   token,
	})
	return
}
