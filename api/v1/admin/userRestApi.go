package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	code2 "gostars/utils/code"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
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

	data, total := models.GetUsers(pageSize, pageNum)

	code := code2.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": code2.GetErrMsg(code),
		},
	)

}

func GetUsersByName(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Param("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := models.GetUsersByUsername(username, pageSize, pageNum)

	code := code2.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": code2.GetErrMsg(code),
		},
	)

}

func EditUser(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := models.CheckUpUser(id, data.Username)
	if code == code2.SUCCESS {
		models.EditUser(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": code2.GetErrMsg(code),
		},
	)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := models.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": code2.GetErrMsg(code),
		},
	)
}
