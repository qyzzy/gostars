package admin

import (
	"github.com/gin-gonic/gin"
	"gostars/middlewares"
)

type LoggerRouterGroup struct {
}

func (s *LoggerRouterGroup) InitLoggerRouter(Router *gin.RouterGroup) {
	loggerRouter := Router.Group("admin").Use(middlewares.JwtToken())
	{
		loggerRouter.GET("logger")
	}
}
