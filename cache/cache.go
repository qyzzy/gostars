package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gostars/utils"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s%s", utils.CacheHost, utils.CachePort))
		},
	}
}
