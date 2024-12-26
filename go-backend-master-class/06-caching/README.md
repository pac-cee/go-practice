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
