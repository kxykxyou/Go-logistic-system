package svc

import (
	"logistic/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        c.Redis.ADDRESS,
		Password:    c.Redis.PASSWORD,
		DB:          c.Redis.DB,
		IdleTimeout: time.Duration(c.Redis.IdleTimeoutSecond) * time.Second,
	})
}
