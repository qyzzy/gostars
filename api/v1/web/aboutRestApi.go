package web

import (
	"github.com/gin-gonic/gin"
	"gostars/utils/code"
	"net/http"
	"strconv"
)

func (userApi *UserApi) GetMe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, errCode := webUserService.GetUserByID(id)

	maps["username"] = data.Username
	maps["role"] = data.Role

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    maps,
			"total":   1,
			"message": code.GetErrMsg(errCode),
		},
	)
}
