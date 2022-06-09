package web

import (
	"github.com/gin-gonic/gin"
	"gostars/models"
	code2 "gostars/utils/code"
	"net/http"
	"strconv"
)

func GetMe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := models.GetUser(id)

	maps["username"] = data.Username
	maps["role"] = data.Role

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": code2.GetErrMsg(code),
		},
	)
}
