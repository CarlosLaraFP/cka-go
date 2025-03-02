package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

/*
Test Type			  Best Practice
-----------------------------------------------
Unit Tests			  Mock Redis
Integration Tests	  Use a Real Redis Instance
*/

// go test ./... -v

// Mock Redis test
func TestRedisConnection(t *testing.T) {
	ctx := context.Background()

	// Start a mock Redis server
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	// Use the mock Redis address
	redisService := NewRedisService(mr.Addr())

	// Set a key
	err = redisService.client.Set(ctx, "testkey", "testvalue", 0).Err()
	assert.NoError(t, err)

	// Get the key
	val, err := redisService.client.Get(ctx, "testkey").Result()
	assert.NoError(t, err)
	assert.Equal(t, "testvalue", val)
}

// Test Redis HTTP Handlers
func TestSetKeyHandler(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	redisService := NewRedisService(mr.Addr())

	req, err := http.NewRequest("POST", "/set?key=test&value=123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redisService.setKey)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Key 'test' set successfully")
}

func TestGetKeyHandler(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	// Create RedisService instance using the mock Redis
	redisService := NewRedisService(mr.Addr())

	// Insert key into Redis using the same client
	err = redisService.client.Set(ctx, "test", "value123", 0).Err()
	assert.NoError(t, err)

	// Initialize a Chi router and bind the route to match real-world usage
	r := chi.NewRouter()
	r.Get("/get/{key}", redisService.getKey)

	// Create HTTP request for the test key
	req, err := http.NewRequest("GET", "/get/test", nil)
	assert.NoError(t, err)

	// Record response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req) // Use the router to process request

	// Assert HTTP response status
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200, got %d", rr.Code)

	// Assert response body
	assert.Contains(t, rr.Body.String(), `"key":"test"`)
	assert.Contains(t, rr.Body.String(), `"value":"value123"`)
}
