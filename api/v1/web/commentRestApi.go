package web

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"gostars/utils/ipsource"
	"net/http"
	"strconv"
)

func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var data models.Comment
	_ = c.ShouldBindJSON(&data)
	articleID, _ := strconv.Atoi(c.Param("id"))
	data.ArticleID = articleID
	data.IPSource = ipsource.OnlineIpInfo(c.Request.RemoteAddr).Region

	errCode := webCommentService.CreateComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"message": code.GetErrMsg(errCode),
	})
}

func (commentApi *CommentApi) GetArticleComments(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	data, errCode := webCommentService.GetArticleComments(articleID)
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    data,
		"message": code.GetErrMsg(errCode),
	})
}

func (commentApi *CommentApi) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	errCode := webCommentService.DeleteComment(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
	})
}
