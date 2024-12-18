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
