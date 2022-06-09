package routes

import (
	"github.com/gin-gonic/gin"
	"gostars/api/v1/web"
	"gostars/utils"
)

var r *gin.Engine

func init() {
	gin.SetMode(utils.AppMode)
	r = gin.New()
	r.Use(gin.Recovery())

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
		//router.GET("user/:id", web.GetUserInfo)
		//router.GET("users", web.GetUsers)
	}

	auth := r.Group("api/v1")
	{
		// 用户模块的路由接口
		auth.GET("admin/users")
		//auth.PUT("user/:id", v1.EditUser)
		//auth.DELETE("user/:id", v1.DeleteUser)
	}

	r.Run(utils.HttpPort)
}
