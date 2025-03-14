package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis"
)

/*
Redis Usage				Uses dependency injection via RedisService
Unit Test Strategy		Uses miniredis mock
Testability				Handlers & Redis logic can be tested separately
Code Structure			Redis logic is encapsulated in RedisService
*/

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
		log.Printf("Failed to set key '%s' in Redis: %v", key, err)
		http.Error(w, fmt.Sprintf("Failed to set key '%s' in Redis: %v", key, err), http.StatusInternalServerError)
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
		log.Printf("Failed to retrieve key '%s' in Redis: %v", key, err)
		http.Error(w, fmt.Sprintf("Failed to retrieve key '%s' in Redis: %v", key, err), http.StatusInternalServerError)
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
	redisService := InitRedis()

	// Setup routes
	r := chi.NewRouter()
	r.Get("/", readRoot)
	r.Get("/health", healthCheck)
	r.Post("/set", redisService.setKey)
	r.Get("/get/{key}", redisService.getKey)
	r.Get("/multiply/{a}/{b}", multiply)

	containerPort := getEnv("CONTAINER_PORT", "8000")

	fmt.Printf("Server running on port %s per Kubernetes manifest\n", containerPort)
	log.Fatal(http.ListenAndServe(":8000", r))
}
