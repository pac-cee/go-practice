# Go Web Basics

This module demonstrates fundamental concepts of web development in Go using the standard `net/http` package.

## Code Explanation

The `main.go` file implements a basic HTTP server with the following components:

### 1. Response Structure
```go
type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```
- Defines a standard API response structure
- Uses JSON tags for proper serialization
- `omitempty` tag skips empty Data field in response
