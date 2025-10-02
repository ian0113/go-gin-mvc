package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/ian0113/go-gin-mvc/infra"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AuthRepository struct {
	ctx         context.Context
	logger      *zap.Logger
	redisClient *redis.Client
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		ctx:         context.TODO(),
		logger:      infra.GetLogger().Named("auth.repository"),
		redisClient: infra.GetRedis(),
	}
}

func (x *AuthRepository) SetRefreshToken(token string, allow bool, expiration time.Duration) error {
	return x.redisClient.Set(x.ctx, "auth:refresh_token:"+token, allow, expiration).Err()
}

func (x *AuthRepository) GetRefreshToken(token string) (bool, error) {
	allowStr, err := x.redisClient.Get(x.ctx, "auth:refresh_token:"+token).Result()
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(allowStr)
}

func (x *AuthRepository) DelRefreshToken(token string) error {
	return x.redisClient.Del(x.ctx, "auth:refresh_token:"+token).Err()
}

func (x *AuthRepository) SetAccessToken(token string, userID uint, expiration time.Duration) error {
	return x.redisClient.Set(x.ctx, "auth:access_token:"+token, userID, expiration).Err()
}

func (x *AuthRepository) GetAccessToken(token string) (uint, error) {
	userIDStr, err := x.redisClient.Get(x.ctx, "auth:access_token:"+token).Result()
	if err != nil {
		return 0, err
	}
	userIDUi64, err := strconv.ParseUint(userIDStr, 10, 32)
	return uint(userIDUi64), err
}

func (x *AuthRepository) DelAccessToken(token string) error {
	return x.redisClient.Del(x.ctx, "auth:access_token:"+token).Err()
}

// func (x *AuthRepository) SetAccessTokenBlacklist(token string, expiration time.Duration) error {
// 	return x.redisClient.Set(x.ctx, "auth:access_token_blacklist:"+token, "1", expiration).Err()
// }

// func (x *AuthRepository) GetAccessTokenBlacklist(token string) (string, error) {
// 	return x.redisClient.Get(x.ctx, "auth:access_token_blacklist:"+token).Result()
// }
