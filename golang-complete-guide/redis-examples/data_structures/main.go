package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	fmt.Println("=== Lists (Queue/Stack) Example ===")
	// Lists - can be used as queues or stacks
	{
		// Push items to list (queue)
		err := rdb.LPush(ctx, "queue", "first", "second", "third").Err()
		if err != nil {
			fmt.Printf("Failed to push to list: %v\n", err)
			return
		}

		// Pop item from queue (FIFO - First In First Out)
		val, err := rdb.RPop(ctx, "queue").Result()
		if err != nil {
			fmt.Printf("Failed to pop from list: %v\n", err)
			return
		}
		fmt.Printf("Popped from queue: %v\n", val)

		// Get all items in list
		items, err := rdb.LRange(ctx, "queue", 0, -1).Result()
		if err != nil {
			fmt.Printf("Failed to get list range: %v\n", err)
			return
		}
		fmt.Printf("Remaining items in queue: %v\n", items)
	}

	fmt.Println("\n=== Sets Example ===")
	// Sets - collection of unique elements
	{
		// Add elements to set
		err := rdb.SAdd(ctx, "fruits", "apple", "banana", "orange").Err()
		if err != nil {
			fmt.Printf("Failed to add to set: %v\n", err)
			return
		}

		// Check if element exists in set
		exists, err := rdb.SIsMember(ctx, "fruits", "apple").Result()
