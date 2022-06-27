package service

var (
	likeService = new(LikeService)
	RmqLikeAdd  *LikeMQ
	RmqLikeDel  *LikeMQ
)

func init() {
	RmqLikeAdd = NewLikeRabbitMQ("like_add")
	go RmqLikeAdd.Consumer()

	RmqLikeDel = NewLikeRabbitMQ("like_del")
	go RmqLikeDel.Consumer()
}
