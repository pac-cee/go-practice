# Go API Design

This module demonstrates how to design and implement a RESTful API in Go using the `gorilla/mux` router, following REST principles and best practices.

## Code Explanation

The `main.go` file implements a complete REST API for product management with the following components:

### 1. Data Models
```go
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    CreatedAt   time.Time `json:"created_at"`
}

type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```
- Clean data structures with JSON tags
- Standardized API response format
- Proper type definitions for all fields

### 2. In-Memory Store
```go
type ProductStore struct {
    sync.RWMutex
    products map[string]Product
}
```
- Thread-safe implementation using `sync.RWMutex`
- Efficient map-based storage
- Demonstrates concurrent access patterns
