package web

import (
	"github.com/gin-gonic/gin"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func (articleApi *ArticleApi) GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, errCode, total := webArticleService.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"total":   total,
		"message": code.GetErrMsg(errCode),
	})
}

func (articleApi *ArticleApi) GetArticlesByTitle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	if len(title) == 0 {
		data, errCode, total := webArticleService.GetArticles(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"data":    data,
			"total":   total,
			"message": code.GetErrMsg(errCode),
		})
		return
	}

	data, errCode, total := webArticleService.GetArticlesByTitle(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"total":   total,
		"message": code.GetErrMsg(errCode),
	})
}

func (articleApi *ArticleApi) GetArticleByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, errCode, total := webArticleService.GetArticlesByCategory(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"total":   total,
		"message": code.GetErrMsg(errCode),
	})
}
