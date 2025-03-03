package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// Context for Redis operations
var ctx = context.Background()

// RedisClient interface for mocking in tests
type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

// RedisService wraps the actual Redis client
type RedisService struct {
	client RedisClient
}

// NewRedisService initializes the Redis client
func NewRedisService(addr string) *RedisService {
	client := redis.NewClient(&redis.Options{Addr: addr})
	return &RedisService{client: client}
}

// Get environment variable with default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// InitRedis initializes Redis with environment variables
func InitRedis() *RedisService {
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	return NewRedisService(redisAddr)
}
