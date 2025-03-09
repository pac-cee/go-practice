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
