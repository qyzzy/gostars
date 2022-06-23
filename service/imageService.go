package service

import (
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type ImageService struct {
}

func (imageService *ImageService) CreateImage(address string, userID int) int {
	var image = new(models.Image)
	image.Address = address
	image.UserID = userID
	err := global.GDb.Table(models.ImageTableName()).
		Create(&image).Error
	if err != nil {
		return code.ErrorImageCreateFailed
	}
	return code.SUCCESS
}
