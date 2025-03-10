# Testing and Debugging in Go 🧪

## 📖 Table of Contents
1. [Unit Testing](#unit-testing)
2. [Table-Driven Tests](#table-driven-tests)
3. [Benchmarking](#benchmarking)
4. [Profiling](#profiling)
5. [Race Detection](#race-detection)
6. [Debugging](#debugging)
7. [Best Practices](#best-practices)
8. [Exercises](#exercises)

## Unit Testing

### Test File Structure
```go
package mypackage_test  // Use separate package for black-box testing

import (
    "testing"
    "github.com/your/package"
)

func TestMyFunction(t *testing.T) {
    // Test implementation
}
```

### Basic Test Function
```go
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

### Test Helper Functions
```go
func TestComplex(t *testing.T) {
    t.Helper()  // Marks this as a helper function
    // Helper implementation
}
```

### Subtests
```go
func TestMath(t *testing.T) {
    t.Run("addition", func(t *testing.T) {
        if got := Add(2, 3); got != 5 {
            t.Errorf("Add(2, 3) = %d; want 5", got)
        }
    })

    t.Run("multiplication", func(t *testing.T) {
        if got := Multiply(2, 3); got != 6 {
            t.Errorf("Multiply(2, 3) = %d; want 6", got)
        }
    })
}
```

## Table-Driven Tests

### Basic Structure
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        x, y int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"mixed", -2, 3, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.x, tt.y)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.x, tt.y, got, tt.want)
            }
        })
    }
}
```

### Testing Errors
```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name    string
        x, y    int
        want    int
        wantErr bool
    }{
        {"valid", 6, 2, 3, false},
        {"zero divisor", 6, 0, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Divide(tt.x, tt.y)
            if (err != nil) != tt.wantErr {
                t.Errorf("Divide(%d, %d) error = %v; wantErr %v",
                    tt.x, tt.y, err, tt.wantErr)
                return
            }
            if !tt.wantErr && got != tt.want {
                t.Errorf("Divide(%d, %d) = %d; want %d",
                    tt.x, tt.y, got, tt.want)
            }
        })
    }
}
```

## Benchmarking

### Basic Benchmark
```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(10)
    }
}
```

### Benchmark with Setup
```go
func BenchmarkComplexOperation(b *testing.B) {
    // Setup code here
    data := makeTestData()
    
    b.ResetTimer()  // Reset timer after setup
    for i := 0; i < b.N; i++ {
        ComplexOperation(data)
    }
}
```

### Parallel Benchmark
```go
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Code to benchmark
        }
    })
}
```

## Profiling

### CPU Profiling
```go
import "runtime/pprof"

func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // Your program here
}
```

### Memory Profiling
```go
import "runtime/pprof"

func main() {
    f, _ := os.Create("mem.prof")
    defer f.Close()
    
    // Your program here
    
    pprof.WriteHeapProfile(f)
}
```

### Using go test for Profiling
```bash
# CPU profile
go test -cpuprofile=cpu.prof -bench .

# Memory profile
go test -memprofile=mem.prof -bench .

# Block profile
go test -blockprofile=block.prof -bench .
```

## Race Detection

### Running Tests with Race Detector
```bash
go test -race ./...
```

### Common Race Conditions
```go
// Data race
var counter int
go func() { counter++ }()
go func() { counter++ }()

// Fix with mutex
var mu sync.Mutex
var counter int
go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

## Debugging

### Using Delve Debugger
```bash
# Install Delve
go install github.com/go-delve/delve/cmd/dlv@latest

