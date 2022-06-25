package admin

import (
	"github.com/gin-gonic/gin"
)

type JwtRouterGroup struct {
}

func (s *JwtRouterGroup) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("")
	{
		jwtRouter.POST("jwt/blacklists", adminApiGroup.AddBlackList)
	}
}
