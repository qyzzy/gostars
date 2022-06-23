package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/utils/code"
	"gostars/utils/upload"
	"net/http"
	"strconv"
)

func (uploadApi *UploadApi) UploadImg(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	userID, _ := strconv.Atoi(c.Query("id"))

	fileSize := fileHeader.Size

	url, errCode := upload.QiNiuUpLoadFile(file, fileSize)

	if errCode == code.SUCCESS {
		createCode := adminImageService.CreateImage(url, userID)
		if createCode != code.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code.GetErrMsg(errCode),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": code.GetErrMsg(errCode),
		"url":     url,
	})
}
