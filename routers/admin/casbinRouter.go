package admin

import (
	"github.com/gin-gonic/gin"
)

type CasbinRouterGroup struct {
}

func (s *CasbinRouterGroup) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("")
	{
		casbinRouter.POST("")
	}
}
