package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
)

type AdminRouterGroup struct {
}

func (s *AdminRouterGroup) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin").Use(middlewares.JwtToken())
	{
		// 用户模块的路由接口
		adminRouter.GET("users", adminApiGroup.GetUsers)
		adminRouter.GET("users/:username", adminApiGroup.GetUsersByName)
		adminRouter.PUT("users/:id", adminApiGroup.EditUser)
		adminRouter.DELETE("users/:id", adminApiGroup.DeleteUser)

		adminRouter.POST("categories", adminApiGroup.CreateCategory)

		adminRouter.POST("articles", adminApiGroup.CreateArticle)
		adminRouter.PUT("articles/:id", adminApiGroup.EditArticle)
		adminRouter.DELETE("articles/:id", adminApiGroup.DeleteArticle)
	}
}
