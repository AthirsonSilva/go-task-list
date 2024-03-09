package database

import (
	"context"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/logger"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Redis", err.Error())
	}

	RedisClient = client
}

func SetRedisObj(username string, token string) {
	RedisClient.Set(context.Background(), username, token, 0)
}

func GetRedisObj(username string) (string, error) {
	return RedisClient.Get(context.Background(), username).Result()
}
