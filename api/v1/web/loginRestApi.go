package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
	"time"
)

func LoginFront(c *gin.Context) {
	var formData models.User
	var errCode int
	_ = c.ShouldBindJSON(&formData)

	formData, errCode = models.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": code.GetErrMsg(errCode),
	})
}

func Login(c *gin.Context) {
	var formData models.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var errCode int

	formData, errCode = models.CheckLogin(formData.Username, formData.Password)

	if errCode == code.SUCCESS {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": code.GetErrMsg(errCode),
			"token":   token,
		})
	}

}

func setToken(c *gin.Context, user models.User) {
	j := middlewares.NewJwt()
	claims := middlewares.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "GoStars",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  code.ERROR,
			"message": code.GetErrMsg(code.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": code.GetErrMsg(200),
		"token":   token,
	})
	return
}
