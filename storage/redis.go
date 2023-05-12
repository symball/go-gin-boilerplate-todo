package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/symball/go-gin-boilerplate/config"
	"log"
)

var redisHandle *redis.Client

func RedisInit(ctx context.Context) {

	redisHandle = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.RedisAddress,
	})

	_, err := redisHandle.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}
