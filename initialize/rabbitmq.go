package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"gostars/global"
	"gostars/utils"
	"log"
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
}
