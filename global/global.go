package global

import (
	"github.com/dutchcoders/goftp"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"gostars/utils/timer"
)

var (
	GDb                *gorm.DB
	GRedis             *redis.Client
	GRouter            *gin.Engine
	GTimer             timer.Timer = timer.NewTimerTask()
	GFileServerConnect *ssh.Client
	GFtp               *goftp.FTP
	GRabbitMQ          *amqp.Connection
)
