package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"gostars/global"
	"gostars/middlewares/rabbitmq"
	"gostars/utils"
	"log"
)

var (
	RmqLikeAdd *rabbitmq.LikeMQ
	RmqLikeDel *rabbitmq.LikeMQ
)

func init() {
	var err error
	global.GRabbitMQ, err = amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		utils.RabbitMQUser,
		utils.RabbitMQPassword,
		utils.RabbitMQAddress,
		utils.RabbitMQPort,
	))

	if err != nil {
		log.Println(err)
	}

	InitLikeRabbitMQ()
}

func InitLikeRabbitMQ() {
	RmqLikeAdd = rabbitmq.NewLikeRabbitMQ("like_add")
	go RmqLikeAdd.Consumer()

	RmqLikeDel = rabbitmq.NewLikeRabbitMQ("like_del")
	go RmqLikeDel.Consumer()
}
