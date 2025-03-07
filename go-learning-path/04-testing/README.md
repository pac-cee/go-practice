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
