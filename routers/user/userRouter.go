package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouterGroup struct {
}

func (s *UserRouterGroup) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("")
	{
		// User info module
		userRouter.POST("users", userApiGroup.Register)
		userRouter.GET("users/:id", userApiGroup.GetMe)

		userRouter.POST("login", userApiGroup.Login)
		userRouter.POST("loginfront", userApiGroup.LoginFront)

		// Article info module
		userRouter.GET("articles", userApiGroup.GetArticles)
		userRouter.GET("articles/title", userApiGroup.GetArticlesByTitle)
		userRouter.GET("articles/:id/comments", userApiGroup.GetArticleComments)

		userRouter.GET("categories", userApiGroup.GetArticles)
		userRouter.GET("categories/:id/articles", userApiGroup.GetArticleByCategory)

		userRouter.POST("comments", userApiGroup.CreateComment)
		userRouter.DELETE("comments/:id", userApiGroup.DeleteComment)
	}
}
