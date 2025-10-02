package infra

import (
	"context"
	"fmt"

	"github.com/ian0113/go-gin-mvc/config"

	"github.com/redis/go-redis/v9"
)

var (
	globalRedisClient *redis.Client
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%d", cfg.Redis.HostName, cfg.Redis.HostPort)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return redisClient

}

func InitRedis(cfg *config.Config) {
	globalRedisClient = NewRedisClient(cfg)
}

func GetRedis() *redis.Client {
	return globalRedisClient
}
