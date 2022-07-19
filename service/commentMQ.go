package service

import "github.com/streadway/amqp"

type CommentMQ struct {
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}
