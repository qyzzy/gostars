package initialize

import (
	"github.com/go-redis/redis/v8"
	"gostars/global"
	"gostars/utils"
)

var (
	redisJwt           = new(RedisJwt)
	redisLikeUserID    = new(RedisLikeUserID)
	redisLikeArticleID = new(RedisLikeArticleID)
)

type RedisServe interface {
	InitRedis() *redis.Client
}

func init() {
	initRedisGroup(
		redisJwt,
		redisLikeUserID,
		redisLikeArticleID,
	)
}

func initRedisGroup(args ...RedisServe) {
	global.GRedisGroup = make([]*redis.Client, 0)
	for _, v := range args {
		global.GRedisGroup = append(global.GRedisGroup, v.InitRedis())
	}
}

type RedisJwt struct{}

func (redisJwt *RedisJwt) InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.CacheHost + utils.CachePort,
		DB:   0,
	})
	return client
}

type RedisLikeUserID struct{}

func (redisLikeUserID *RedisLikeUserID) InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.CacheHost + utils.CachePort,
		DB:   1,
	})
	return client
}

type RedisLikeArticleID struct{}

func (redisLikeArticleID *RedisLikeArticleID) InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.CacheHost + utils.CachePort,
		DB:   2,
	})
	return client
}
