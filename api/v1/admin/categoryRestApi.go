package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
)

func (categoryApi *CategoryApi) CreateCategory(c *gin.Context) {
	var data models.Category
	_ = c.ShouldBindJSON(&data)
	errCode := adminCategoryService.CheckCategory(data.CategoryName)
	fmt.Println(data.CategoryName)
	if errCode == code.SUCCESS {
		adminCategoryService.CreateCategory(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    data,
			"message": code.GetErrMsg(errCode),
		},
	)
}
