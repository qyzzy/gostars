package global

import (
	"github.com/dutchcoders/goftp"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/importcjj/sensitive"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"gostars/utils/timer"
	"sync"
)

var (
	GWG                sync.WaitGroup
	GDb                *gorm.DB
	GRouter            *gin.Engine
	GTimer             timer.Timer = timer.NewTimerTask()
	GFileServerConnect *ssh.Client
	GFtp               *goftp.FTP
	GRabbitMQ          *amqp.Connection
	GRedisGroup        []*redis.Client
	GFilter            *sensitive.Filter
)
