package user

import (
	"github.com/gin-gonic/gin"
	"gostars/api/v1/web"
)

type UserRouterGroup struct {
}

func (s *UserRouterGroup) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("")
	{
		// User info module
		userRouter.POST("users", web.Register)
		userRouter.GET("users/:id", web.GetMe)

		userRouter.POST("login", web.Login)
		userRouter.POST("loginfront", web.LoginFront)

		// Article info module
		userRouter.GET("articles", web.GetArticles)
		userRouter.GET("articles/title", web.GetArticlesByTitle)
		userRouter.GET("articles/:id/comments", web.GetArticleComments)

		userRouter.GET("categories")
		userRouter.GET("categories/:id/articles", web.GetArticleByCategory)

		userRouter.POST("comments", web.CreateComment)
		userRouter.DELETE("comments/:id", web.DeleteComment)
	}
}
