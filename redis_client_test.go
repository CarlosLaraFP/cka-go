package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	initRedis()
	err := redisClient.Set(ctx, "testkey", "testvalue", 0).Err()
	assert.NoError(t, err)

	val, err := redisClient.Get(ctx, "testkey").Result()
	assert.NoError(t, err)
	assert.Equal(t, "testvalue", val)
}
