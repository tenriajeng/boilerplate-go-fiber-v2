package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func NewRedis(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddr(),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,

		// Connection pool settings
		PoolSize:     10, // Number of connections in pool
		MinIdleConns: 5,  // Minimum idle connections
		MaxRetries:   3,  // Maximum retry attempts

		// Timeouts
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Println("Warning: Failed to connect to Redis:", err)
		log.Println("Application will continue without Redis")
		return nil
	}

	log.Println("Redis connected successfully")
	RedisClient = rdb
	return rdb
}

func GetRedis() *redis.Client {
	return RedisClient
}
