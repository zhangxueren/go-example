package helper

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func initRedis(ctx context.Context) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	Logger.Info("Connected to Redis successfully!")
}
