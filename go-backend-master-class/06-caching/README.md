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
