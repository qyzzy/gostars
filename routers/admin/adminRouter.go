package admin

import (
	"github.com/gin-gonic/gin"
)

type AdminRouterGroup struct {
}

func (s *AdminRouterGroup) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("")
	{
		// user module
		adminRouter.GET("users", adminApiGroup.GetUsers)
		adminRouter.GET("users/:username", adminApiGroup.GetUsersByName)
		adminRouter.PUT("users/:id", adminApiGroup.EditUser)
		adminRouter.DELETE("users/:id", adminApiGroup.DeleteUser)

		adminRouter.POST("categories", adminApiGroup.CreateCategory)

		adminRouter.GET("articles")
		adminRouter.POST("articles", adminApiGroup.CreateArticle)
		adminRouter.PUT("articles/:id", adminApiGroup.EditArticle)
		adminRouter.DELETE("articles/:id", adminApiGroup.DeleteArticle)
	}
}
