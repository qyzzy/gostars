package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/api/v1/admin"
	"gostars/middlewares"
)

type AdminRouterGroup struct {
}

func (s *AdminRouterGroup) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin").Use(middlewares.JwtToken())
	{
		// 用户模块的路由接口
		adminRouter.GET("users", admin.GetUsers)
		adminRouter.GET("users/:username", admin.GetUsersByName)
		adminRouter.PUT("users/:id", admin.EditUser)
		adminRouter.DELETE("users/:id", admin.DeleteUser)

		adminRouter.POST("categories", admin.CreateCategory)

		adminRouter.POST("articles", admin.CreateArticle)
		adminRouter.PUT("articles/:id", admin.EditArticle)
		adminRouter.DELETE("articles/:id", admin.DeleteArticle)
	}
}
