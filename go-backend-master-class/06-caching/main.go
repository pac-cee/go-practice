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

// NewInMemoryCache creates a new in-memory cache
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]string),
	}
}

// Redis Cache implementation
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// In-Memory Cache implementation
func (c *InMemoryCache) Get(ctx context.Context, key string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if val, ok := c.data[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("key not found")
}

func (c *InMemoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c.Lock()
	defer c.Unlock()
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.data[key] = string(str)
