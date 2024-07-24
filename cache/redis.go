package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/saas0503/factory-api/config"
	"log"
)

var RedisInstance *redis.Client

func ConnectRedis(cfg config.ApiConfig) {
	RedisInstance = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisUrl,
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	result := RedisInstance.Ping(ctx)
	if result.Err() != nil {
		log.Fatalf("redis connect err: %v", result.Err())
	}
}
