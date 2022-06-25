package admin

import (
	"github.com/gin-gonic/gin"
)

type UploadRouterGroup struct {
}

func (s *UploadRouterGroup) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouter := Router.Group("")
	{
		uploadRouter.POST("upload/images/", adminApiGroup.UploadImg)
	}
}
