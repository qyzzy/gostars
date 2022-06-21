package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
)

type JwtRouterGroup struct {
}

func (s *JwtRouterGroup) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("admin").Use(middlewares.JwtToken())
	{
		jwtRouter.POST("jwt/blacklists", adminApiGroup.AddBlackList)
	}
}
