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
		userRouter.GET("users/:id/favorite", userApiGroup.GetFavoriteList)

		userRouter.POST("login", userApiGroup.Login)
		userRouter.POST("loginfront", userApiGroup.LoginFront)

		// Article info module
		userRouter.GET("articles", userApiGroup.GetArticles)
		userRouter.GET("articles/title", userApiGroup.GetArticlesByTitle)
		userRouter.GET("articles/:id/comments", userApiGroup.GetArticleComments)
		userRouter.POST("articles/:id/favorite", userApiGroup.FavoriteAction)
		userRouter.POST("articles/:id/comments", userApiGroup.CreateComment)

		// Category info module
		userRouter.GET("categories", userApiGroup.GetArticles)
		userRouter.GET("categories/:id/articles", userApiGroup.GetArticleByCategory)

		// Comment info module
		userRouter.DELETE("comments/:id", userApiGroup.DeleteComment)
	}
}
