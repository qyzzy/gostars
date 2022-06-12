package routes

import (
	"github.com/gin-gonic/gin"
	"gostars/api/v1/admin"
	"gostars/api/v1/web"
	"gostars/middlewares"
	"gostars/utils"
)

var r *gin.Engine

func init() {
	gin.SetMode(utils.AppMode)
	r = gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	r.Use(middlewares.Logger())

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	router := r.Group("api/v1")
	{
		// User info module
		router.POST("users", web.Register)
		router.GET("users/:id", web.GetMe)
		//router.GET("users", web.GetUsers)

		router.POST("loginfront", web.LoginFront)

		// Article info module
		router.GET("articles", web.GetArticles)
		router.GET("articles/:title", web.GetArticlesByTitle)
		//router.GET("article/info/:id", v1.GetArtInfo)

		router.GET("categories")
		router.GET("categories/:id/articles", web.GetArticleByCategory)
	}

	auth := r.Group("api/v1/admin")
	auth.Use(middlewares.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("users", admin.GetUsers)
		auth.GET("users/:username", admin.GetUsersByName)
		auth.PUT("users/:id", admin.EditUser)
		auth.DELETE("users/:id", admin.DeleteUser)

		auth.POST("categories", admin.CreateCategory)

		auth.POST("articles", admin.CreateArticle)
		auth.PUT("articles/:id", admin.EditArticle)
		auth.DELETE("articles/:id", admin.DeleteArticle)
	}

	_ = r.Run(utils.HttpPort)
}
