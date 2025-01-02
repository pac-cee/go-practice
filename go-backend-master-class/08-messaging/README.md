# Go Messaging System

This module demonstrates how to implement a real-time messaging system in Go using WebSockets and a pub/sub pattern.

## Code Explanation

The `main.go` file implements a complete messaging system with the following components:

### 1. Data Models
```go
type Message struct {
    ID        string      `json:"id"`
    Topic     string      `json:"topic"`
    Data      interface{} `json:"data"`
    Timestamp time.Time   `json:"timestamp"`
}

type Subscription struct {
    Topic    string
