package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
)

var (
	RedisClient *redis.Client
)

func Init() {
	config := env.LoadConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", config.RedisPort),
		Password: "", // If your Redis requires a password, you can add it to your env config and use it here
		DB:       0,  // Default DB
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
