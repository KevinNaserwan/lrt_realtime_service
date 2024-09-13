package redis

import (
	"context"
	"fmt"
	"log"

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
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Redis client successfully connected.")
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
