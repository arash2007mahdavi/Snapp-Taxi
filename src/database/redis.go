package database

import (
	"fmt"
	"snapp/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB: 0,
		DialTimeout: cfg.Redis.DialTimeout * time.Second,
		ReadTimeout: cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize: cfg.Redis.PoolSize,
		PoolTimeout: cfg.Redis.PoolTimeout * time.Second,
	})
}

func GetRedis() *redis.Client{
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}