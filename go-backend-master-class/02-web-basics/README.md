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

### 2. Middleware Implementation
```go
func loggingMiddleware(next http.Handler) http.Handler
```
- Demonstrates the middleware pattern in Go
- Logs request start and completion times
- Shows how to chain HTTP handlers
- Uses `time.Since()` for duration measurement

### 3. JSON Response Helper
```go
func respondWithJSON(w http.ResponseWriter, code int, payload interface{})
```
- Centralizes JSON response handling
- Sets proper Content-Type header
- Handles HTTP status codes
- Marshals payload to JSON

### 4. Route Handlers
```go
func HomeHandler(w http.ResponseWriter, r *http.Request)
func HealthCheckHandler(w http.ResponseWriter, r *http.Request)
```
- `HomeHandler`: Basic route handling with path validation
- `HealthCheckHandler`: System health status endpoint
- Demonstrates proper error handling
- Shows how to structure API responses

### 5. Server Configuration
```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      handler,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
}
```
