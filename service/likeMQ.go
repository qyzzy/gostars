package service

import (
	"fmt"
	"github.com/streadway/amqp"
	"gostars/global"
	"gostars/models"
	"gostars/utils"
	"gostars/utils/code"
	"log"
	"strconv"
	"strings"
)

type LikeMQ struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

func NewLikeRabbitMQ(queueName string) *LikeMQ {
	likeMQ := &LikeMQ{
		queueName: queueName,
	}
	channel, err := global.GRabbitMQ.Channel()
	if err != nil {
		log.Println(err)
	}
	likeMQ.channel = channel
	return likeMQ
}

func (likeMQ *LikeMQ) Publish(message string) {
	_, err := likeMQ.channel.QueueDeclare(
		likeMQ.queueName,
		// durable
		false,
		// autoDelete
		false,
		// exclusive
		false,
		// wait or not
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	// publish message
	err = likeMQ.channel.Publish(
		likeMQ.exchange,
		likeMQ.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Println(err)
	}
}

func (likeMQ *LikeMQ) Consumer() {
	_, err := likeMQ.channel.QueueDeclare(
		likeMQ.queueName,
		// durable
		false,
		// autoDelete
		false,
		// exclusive
		false,
		// wait or not
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	// consume message
	messages, err := likeMQ.channel.Consume(
		likeMQ.exchange,
		likeMQ.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	durable := make(chan bool)
	switch likeMQ.queueName {
	case "like_add":
		go likeMQ.consumerLikeAdd(messages)
	case "like_del":
		go likeMQ.consumerLikeDel(messages)
	}
	<-durable
}

func (likeMQ *LikeMQ) consumerLikeAdd(messages <-chan amqp.Delivery) {
	for d := range messages {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userID, _ := strconv.Atoi(params[0])
		articleID, _ := strconv.Atoi(params[1])

		for i := 0; i < utils.Attempts; i++ {
			// default : no problem
			flag := false
			var likeData models.Like
			_, errCode := likeService.GetLikeInfo(articleID, userID)
			if errCode == code.ERROR {
				log.Println(errCode)
				// something wrong
				flag = true
			} else {
				if errCode == code.ErrorDataNotFound {
					likeData.ArticleID = articleID
					likeData.UserID = userID
					likeData.Cancel = utils.IsLike

					errCode = likeService.CreateLike(&likeData)
					if errCode != code.SUCCESS {
						log.Println(errCode)
						flag = true
					}
				} else {
					errCode = likeService.UpdateLike(articleID, userID, utils.IsLike)
					if errCode != code.SUCCESS {
						log.Println(errCode)
						flag = true
					}
				}
			}
			// end the loop when it executes normally
			if !flag {
				break
			}
		}
	}
}

func (likeMQ *LikeMQ) consumerLikeDel(messages <-chan amqp.Delivery) {
	for d := range messages {
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userID, _ := strconv.Atoi(params[0])
		articleID, _ := strconv.Atoi(params[1])

		for i := 0; i < utils.Attempts; i++ {
			flag := true
			_, errCode := likeService.GetLikeInfo(articleID, userID)
			if errCode == code.ERROR {
				log.Println(errCode)
				// something wrong
				flag = true
			} else {
				if errCode == code.ErrorDataNotFound {
					log.Println("MySQL and RabbitMQ wrong!")
				} else {
					errCode = likeService.UpdateLike(articleID, userID, utils.UnLike)
					if errCode != code.SUCCESS {
						log.Println(errCode)
						flag = true
					}
				}
			}
			if !flag {
				break
			}
		}
	}
}
