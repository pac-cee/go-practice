package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// Product represents a product in our system
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Source  string      `json:"source,omitempty"` // "cache" or "database"
}

// Cache interface defines our caching behavior
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

// RedisCache implements Cache interface using Redis
type RedisCache struct {
	client *redis.Client
}

// InMemoryCache implements Cache interface using a map
type InMemoryCache struct {
	sync.RWMutex
	data map[string]string
}

// NewRedisCache creates a new Redis cache
func NewRedisCache(addr string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{client: client}
}

