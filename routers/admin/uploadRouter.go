package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
)

type UploadRouterGroup struct {
}

func (s *UploadRouterGroup) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouter := Router.Group("admin").Use(middlewares.JwtToken())
	{
		uploadRouter.POST("upload/img/", adminApiGroup.UploadImg)
	}
}
