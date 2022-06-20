package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func (userApi *UserApi) GetUsers(c *gin.Context) {
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

	data, total := adminUserService.GetUsers(pageSize, pageNum)

	errCode := code.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    data,
			"total":   total,
			"message": code.GetErrMsg(errCode),
		},
	)

}

func (userApi *UserApi) GetUsersByName(c *gin.Context) {
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

	data, total := adminUserService.GetUsersByUsername(username, pageSize, pageNum)

	errCode := code.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    data,
			"total":   total,
			"message": code.GetErrMsg(errCode),
		},
	)

}

func (userApi *UserApi) ChangeUserPassword(c *gin.Context) {

}

func (userApi *UserApi) EditUser(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	//code := models.CheckUpUser(id, data.Username)
	//if code == code2.SUCCESS {
	//	models.EditUser(id, &data)
	//}

	errCode := adminUserService.EditUser(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		},
	)
}

func (userApi *UserApi) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	errCode := adminUserService.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": code.GetErrMsg(errCode),
		},
	)
}
