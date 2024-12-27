# Go Caching Implementation

This module demonstrates how to implement caching in Go using both Redis and in-memory caching strategies.

## Code Explanation

The `main.go` file implements a complete caching system with the following components:

### 1. Data Models
```go
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Price       float64   `json:"price"`
    LastUpdated time.Time `json:"last_updated"`
}

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Source  string      `json:"source,omitempty"` // "cache" or "database"
}
```
- Product model with cache-relevant fields
- API response with cache source tracking

### 2. Cache Interface
```go
type Cache interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Delete(ctx context.Context, key string) error
}
```
- Generic cache interface
- Context-aware operations
- Support for expiration

### 3. Redis Cache Implementation
```go
type RedisCache struct {
    client *redis.Client
}
```
- Redis client integration
- Implements Cache interface
- Handles distributed caching
- Supports TTL

### 4. In-Memory Cache Implementation
```go
type InMemoryCache struct {
    sync.RWMutex
    data map[string]string
}
```
- Thread-safe implementation
- Map-based storage
- Local memory caching
- Lock-based synchronization

### 5. API Integration
```go
type API struct {
    cache Cache
    store map[string]Product
    mutex sync.RWMutex
    ctx   context.Context
}
```
- Cache-aware API design
- Strategy pattern for cache selection
- Thread-safe operations
- Context propagation

