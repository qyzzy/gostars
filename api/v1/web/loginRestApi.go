package web

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	code2 "gostars/utils/code"
	"net/http"
)

func LoginFront(c *gin.Context) {
	var formData models.User
	var code int
	_ = c.ShouldBindJSON(&formData)

	formData, code = models.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": code2.GetErrMsg(code),
	})
}
