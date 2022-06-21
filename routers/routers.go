package routers

import (
	"github.com/gin-gonic/gin"
	"gostars/global"
	"gostars/middlewares"
	"gostars/utils"
	"time"
)

func init() {
	initRouter()
}

func initRouter() {
	gin.SetMode(utils.AppMode)
	global.GRouter = gin.New()

	global.GRouter.Use(gin.Recovery())
	global.GRouter.Use(middlewares.Cors())
	global.GRouter.Use(middlewares.Logger())
	global.GRouter.Use(middlewares.RateLimit(time.Second, 100, 100))

	enterGroup := &RouterGroup{}
	adminRouter := enterGroup.Admin
	userRouter := enterGroup.User

	publicGroup := global.GRouter.Group("api/v1")
	{
		userRouter.InitUserRouter(publicGroup)
	}

	privateGroup := global.GRouter.Group("api/v1")
	{
		adminRouter.InitAdminRouter(privateGroup)
		adminRouter.InitLoggerRouter(privateGroup)
		adminRouter.InitJwtRouter(privateGroup)
	}

	_ = global.GRouter.Run(utils.HttpPort)
}
