package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Param("pagesize"))
	pageNum, _ := strconv.Atoi(c.Param("pagenum"))
	username := c.Query("username")

}
