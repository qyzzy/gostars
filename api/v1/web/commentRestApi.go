package web

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func CreateComment(c *gin.Context) {
	var data models.Comment
	_ = c.ShouldBindJSON(&data)

	errCode := models.CreateComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"message": code.GetErrMsg(errCode),
	})
}

func GetArticleComments(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	data, errCode := models.GetArticleComments(articleID)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"message": code.GetErrMsg(errCode),
	})
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	errCode := models.DeleteComment(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
	})
}
