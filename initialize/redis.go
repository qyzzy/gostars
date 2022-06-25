package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gostars/global"
	"gostars/utils"
	"log"
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr: utils.CacheHost + utils.CachePort,
		DB:   0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
	}
	global.GRedis = client
}
