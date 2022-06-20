package web

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"gostars/utils/validator"
	"net/http"
)

func (userApi *UserApi) Register(c *gin.Context) {
	var data models.User
	var msg string
	var validCode int

	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)

	if validCode != code.SUCCESS {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	errCode := webUserService.Register(&data)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		},
	)
}
