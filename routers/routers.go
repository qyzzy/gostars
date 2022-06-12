package routers

import (
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
	"gostars/utils"
	"time"
)

var r *gin.Engine

func init() {
	initRouter()
}

func initRouter() {
	gin.SetMode(utils.AppMode)
	r = gin.New()

	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	r.Use(middlewares.Logger())
	r.Use(middlewares.RateLimit(time.Second, 100, 100))

	enterGroup := &RouterGroup{}
	adminRouter := enterGroup.Admin
	userRouter := enterGroup.User

	publicGroup := r.Group("api/v1")
	{
		userRouter.InitUserRouter(publicGroup)
	}

	privateGroup := r.Group("api/v1")
	{
		adminRouter.InitAdminRouter(privateGroup)
	}

	_ = r.Run(utils.HttpPort)
}
