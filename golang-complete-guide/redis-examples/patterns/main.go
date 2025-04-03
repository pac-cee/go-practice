package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Caching example
func cacheExample(ctx context.Context, rdb *redis.Client) {
	cacheKey := "user:profile:123"
	
	// Try to get data from cache
	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss - get data from "database" and cache it
		userData := `{"id": 123, "name": "John Doe", "email": "john@example.com"}`
		err = rdb.Set(ctx, cacheKey, userData, 1*time.Hour).Err()
