# Concurrency in Go 🔄

## 📖 Table of Contents
1. [Goroutines](#goroutines)
2. [Channels](#channels)
3. [Select Statement](#select-statement)
4. [Sync Package](#sync-package)
5. [Concurrency Patterns](#concurrency-patterns)
6. [Best Practices](#best-practices)
7. [Exercises](#exercises)

## Goroutines

### What are Goroutines?
- Lightweight threads managed by Go runtime
- Much smaller stack size than OS threads
- Multiplexed onto OS threads
- Non-blocking by default

### Basic Usage
```go
func main() {
    go myFunction()    // Start a new goroutine
    go func() {        // Anonymous function as goroutine
        // Do something
    }()
}
```

### Key Points
- Goroutines are cheap and lightweight
- Go runtime handles scheduling
- No direct communication between goroutines (use channels)
- Main function runs in the main goroutine

## Channels

### Channel Basics
```go
ch := make(chan int)        // Unbuffered channel
ch := make(chan int, 100)   // Buffered channel with capacity 100

// Send and receive
ch <- value    // Send value
value := <-ch  // Receive value
```

### Channel Types
1. **Unbuffered Channels**
   - Synchronous communication
   - Sender blocks until receiver is ready
   - Receiver blocks until sender sends

2. **Buffered Channels**
   - Asynchronous up to buffer size
   - Sender blocks only when buffer is full
   - Receiver blocks only when buffer is empty

3. **Directional Channels**
```go
chan<- int   // Send-only channel
<-chan int   // Receive-only channel
```

### Channel Operations
```go
// Close channel
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    // Channel is closed
}

// Range over channel
for value := range ch {
    // Process value
}
```

## Select Statement

### Basic Select
```go
select {
case v1 := <-ch1:
    // Handle value from ch1
case v2 := <-ch2:
    // Handle value from ch2
case ch3 <- x:
    // Send x to ch3
default:
    // Optional default case
}
```

### Common Patterns
1. **Timeout**
```go
select {
case res := <-ch:
    // Handle result
case <-time.After(1 * time.Second):
    // Handle timeout
}
```

2. **Non-blocking Operations**
```go
select {
case ch <- x:
    // Sent successfully
default:
    // Would block
}
```

## Sync Package

### Mutex
```go
var mu sync.Mutex
mu.Lock()
// Critical section
mu.Unlock()
```

### RWMutex
```go
var mu sync.RWMutex
mu.RLock()    // Multiple readers
// Read operations
mu.RUnlock()

mu.Lock()     // Single writer
// Write operations
mu.Unlock()
```

### WaitGroup
```go
var wg sync.WaitGroup
wg.Add(n)     // Add n goroutines
go func() {
    defer wg.Done()
    // Work
}()
wg.Wait()     // Wait for all to finish
```

### Once
