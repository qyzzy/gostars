package redis

import (
	"context"
	"gostars/global"
	"gostars/utils/code"
)

type redisOptions struct {
}

func (o *redisOptions) deleteKey(key string) int {
	err := global.GRedis.Del(context.Background(), key).Err()
	if err != nil {
		return code.ErrorRedisDeleteFailed
	}
	return code.SUCCESS
}
