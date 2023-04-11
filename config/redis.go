package config

import (
	"os"

	"github.com/go-redis/redis"
)

func RedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
	})

	return client

}
