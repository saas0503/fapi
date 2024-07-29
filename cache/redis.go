package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/saas0503/factory-api/config"
)

var RedisInstance *redis.Client

func ConnectRedis(cfg config.Config) {
	RedisInstance = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})

	ctx := context.Background()
	result := RedisInstance.Ping(ctx)
	if result.Err() != nil {
		log.Fatalf("redis connect err: %v", result.Err())
	}
}
