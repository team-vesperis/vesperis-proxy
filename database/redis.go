package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/team-vesperis/vesperis-proxy/config"
)

var client *redis.Client

func initializeRedis() {
	opt, urlError := redis.ParseURL(config.GetRedisUrl())
	if urlError != nil {
		logger.Panic("Error parsing url in the Redis Database. - ", urlError)
	}

	client = redis.NewClient(opt)

	setError := client.Set(context.Background(), "ping", "pong", 0).Err()
	if setError != nil {
		logger.Panic("Error sending value to the Redis Database. - ", setError)
	}

	value, getError := client.Get(context.Background(), "ping").Result()
	if getError != nil {
		logger.Panic("Error retrieving value from the Redis Database. - ", getError)
	}

	logger.Info("Checking connection with Redis Database: " + value)
	logger.Info("Successfully initialized the Redis Database.")
}

func getRedisClient() *redis.Client {
	if client == nil {
		logger.Error("Redis client not found.")
	}
	return client
}

func closeRedis() {
	if client != nil {
		client.Close()
	}
}
