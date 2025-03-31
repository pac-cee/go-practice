package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	ctx := context.Background()

	// Ping the Redis server to check connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return
	}
	fmt.Println("Connected to Redis!")

	// Basic SET operation
	err := rdb.Set(ctx, "mykey", "Hello from Redis!", 0).Err()
	if err != nil {
		fmt.Printf("Failed to set key: %v\n", err)
		return
	}

	// Basic GET operation
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		fmt.Printf("Failed to get key: %v\n", err)
		return
