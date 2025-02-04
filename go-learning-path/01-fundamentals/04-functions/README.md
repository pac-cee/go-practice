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
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

### Pointer Receivers
```go
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}
```

### Value vs Pointer Receivers
```go
// Value receiver - gets a copy
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Pointer receiver - can modify the original
func (r *Rectangle) SetWidth(w float64) {
    r.width = w
}
```

## Interfaces

### Interface Declaration
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Implementing Interfaces
```go
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

### Empty Interface
```go
func PrintAnything(v interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", v, v)
}
```

## Function Types

### Function as a Type
```go
type Operation func(a, b int) int

func Calculate(a, b int, op Operation) int {
    return op(a, b)
}
```

### Anonymous Functions
```go
func main() {
    sum := func(a, b int) int {
        return a + b
    }
    
    result := sum(3, 4)
}
```

### Closures
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Error Handling

### Error Return Pattern
```go
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    return nil
}
```

### Custom Errors
```go
type ValidationError struct {
    Field string
    Error string
}
