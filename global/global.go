package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gostars/utils/timer"
)

var (
	GDb     *gorm.DB
	GRedis  *redis.Client
	GRouter *gin.Engine
	GTimer  timer.Timer = timer.NewTimerTask()
)
