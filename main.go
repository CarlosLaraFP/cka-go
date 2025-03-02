package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

/*
Redis Usage				Uses dependency injection via RedisService
Unit Test Strategy		Uses miniredis mock
Testability				Handlers & Redis logic can be tested separately
Code Structure			Redis logic is encapsulated in RedisService
*/

// Context for Redis operations
var ctx = context.Background()

// RedisClient interface to allow mocking in tests
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

// Root handler
func readRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Go running with Redis!"})
}

// Health check handler
func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Set key in Redis
func (rs *RedisService) setKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if key == "" || value == "" {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
		return
	}

	err := rs.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		http.Error(w, "Failed to set key", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Key '%s' set successfully", key)})
}

// Get key from Redis
func (rs *RedisService) getKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	value, err := rs.client.Get(ctx, key).Result()
	if err == redis.Nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve key", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
}

// Multiply two numbers asynchronously
func multiply(w http.ResponseWriter, r *http.Request) {
	aStr := chi.URLParam(r, "a")
	bStr := chi.URLParam(r, "b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Simulate async operation
	time.Sleep(100 * time.Millisecond)

	json.NewEncoder(w).Encode(map[string]int{"result": a * b})
}

func main() {
	// Initialize Redis connection
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisService := NewRedisService(redisAddr)

	r := chi.NewRouter()
	r.Get("/", readRoot)
	r.Get("/health", healthCheck)
	r.Post("/set", redisService.setKey)
	r.Get("/get/{key}", redisService.getKey)
	r.Get("/multiply/{a}/{b}", multiply)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
