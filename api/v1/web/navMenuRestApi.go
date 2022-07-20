package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (navMenuApi *NavMenuApi) GetLevel0NavMenus(c *gin.Context) {
	level0NavMenus, errCode := webLevel0NavMenuService.GetLevel0NavMenus()
	c.JSON(http.StatusOK, gin.H{
		"status":         errCode,
		"level0NavMenus": level0NavMenus,
	})
	return
}

func (navMenuApi *NavMenuApi) GetLevel1NavMenus(c *gin.Context) {
	level1NavMenus, errCode := webLevel1NavMenuService.GetLevel1NavMenus()
	c.JSON(http.StatusOK, gin.H{
		"status":         errCode,
		"level1NavMenus": level1NavMenus,
	})
	return
}
