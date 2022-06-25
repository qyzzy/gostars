package admin

import (
	"github.com/gin-gonic/gin"
)

type LoggerRouterGroup struct {
}

func (s *LoggerRouterGroup) InitLoggerRouter(Router *gin.RouterGroup) {
	loggerRouter := Router.Group("")
	{
		loggerRouter.GET("logger")
	}
}
