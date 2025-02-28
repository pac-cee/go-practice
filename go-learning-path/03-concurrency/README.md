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
