# Functions and Methods in Go 🔧

## 📖 Table of Contents
1. [Function Basics](#function-basics)
2. [Methods](#methods)
3. [Interfaces](#interfaces)
4. [Function Types](#function-types)
5. [Error Handling](#error-handling)
6. [Best Practices](#best-practices)
7. [Exercises](#exercises)

## Function Basics

### Basic Function Declaration
```go
func functionName(param1 Type1, param2 Type2) ReturnType {
    // function body
    return value
}
```

### Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Named Return Values
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // naked return
}
```

### Variadic Functions
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Methods

### Method Declaration
```go
type Rectangle struct {
    width, height float64
