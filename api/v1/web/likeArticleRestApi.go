package web

import (
	"github.com/gin-gonic/gin"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func (likeArticleApi *LikeArticleApi) FavoriteAction(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := strconv.Atoi(c.Query("id"))
	actionType, _ := strconv.Atoi(c.Query("actiontype"))

	errCode := webFavoriteService.FavoriteAction(userID, articleID, actionType)

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
	})
}
