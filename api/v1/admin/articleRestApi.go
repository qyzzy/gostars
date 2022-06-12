package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	code2 "gostars/utils/code"
	"net/http"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	var data models.Article
	_ = c.ShouldBindJSON(&data)

	code := models.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": code2.GetErrMsg(code),
	})
}

func EditArticle(c *gin.Context) {
	var data models.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := models.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": code2.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := models.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": code2.GetErrMsg(code),
	})
}
