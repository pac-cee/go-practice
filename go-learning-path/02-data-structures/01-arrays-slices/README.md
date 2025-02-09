# Arrays and Slices in Go 📚

## 📖 Table of Contents
1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Operations](#operations)
4. [Memory Management](#memory-management)
5. [Best Practices](#best-practices)
6. [Exercises](#exercises)

## Arrays

### Array Declaration
```go
// Fixed-size array
var numbers [5]int

// Array with initialization
scores := [3]int{90, 95, 100}

// Array with size inference
names := [...]string{"Alice", "Bob", "Charlie"}
```

### Array Properties
- Fixed length
- Zero-based indexing
- Contiguous memory
- Pass by value (copying)

### Multi-dimensional Arrays
```go
// 2D array
var matrix [3][3]int

// Initialize 2D array
grid := [2][3]int{
    {1, 2, 3},
