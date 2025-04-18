package database

import (
	"context"
	"encoding/json"
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

func Set[T any](redis *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redis.Set(context.Background(), key, v, duration*time.Second).Err()
}

func Get[T any](redis *redis.Client, key string) (T, error) {
	var dest = *new(T)
	res, err := redis.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(res), dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}