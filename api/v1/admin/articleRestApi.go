package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	var data models.Article
	_ = c.ShouldBindJSON(&data)

	errCode := adminArticleService.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"message": code.GetErrMsg(errCode),
	})
}

func (articleApi *ArticleApi) EditArticle(c *gin.Context) {
	var data models.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	errCode := adminArticleService.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
	})
}

func (articleApi *ArticleApi) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	errCode := adminArticleService.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
	})
}
