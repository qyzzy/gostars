package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostars/models"
	code2 "gostars/utils/code"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	var data models.Category
	_ = c.ShouldBindJSON(&data)
	code := models.CheckCategory(data.CategoryName)
	fmt.Println(data.CategoryName)
	if code == code2.SUCCESS {
		models.CreateCategory(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": code2.GetErrMsg(code),
		},
	)
}
