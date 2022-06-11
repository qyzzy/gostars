package routes

import (
	"github.com/gin-gonic/gin"
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
		router.POST("user/add", web.Register)
		router.GET("user/:id", web.GetMe)
		//router.GET("users", web.GetUsers)

		router.POST("loginfront", web.LoginFront)

		// Article info module
		router.GET("article", web.GetArticles)
		router.GET("article/:title", web.GetArticlesByTitle)
		router.GET("article/category/:id", web.GetArticleByCategory)
		//router.GET("article/info/:id", v1.GetArtInfo)
	}

	auth := r.Group("api/v1")
	auth.Use(middlewares.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users")
		//auth.PUT("user/:id", v1.EditUser)
		//auth.DELETE("user/:id", v1.DeleteUser)
	}

	r.Run(utils.HttpPort)
}
