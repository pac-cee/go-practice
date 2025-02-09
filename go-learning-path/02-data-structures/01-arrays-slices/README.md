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
    {4, 5, 6},
}
```

## Slices

### Slice Declaration
```go
// Empty slice
var numbers []int

// Slice with initialization
scores := []int{90, 95, 100}

// Make with length and capacity
buffer := make([]byte, 5, 10)
```

### Slice Properties
- Dynamic length
- Reference type
- Built on arrays
- Three components:
  - Pointer to array
  - Length
  - Capacity

### Slice Operations

#### Append
```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)
numbers = append(numbers, 5, 6, 7)
```

#### Slicing
```go
// Slice syntax: slice[startIndex:endIndex]
numbers := []int{0, 1, 2, 3, 4, 5}
subset := numbers[2:4]  // [2, 3]
```

#### Copy
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

## Operations

### Iteration
```go
// Using range
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Using traditional for loop
for i := 0; i < len(slice); i++ {
    fmt.Printf("Value at %d: %d\n", i, slice[i])
}
```

### Sorting
```go
// Sort integers
numbers := []int{3, 1, 4, 1, 5, 9}
sort.Ints(numbers)

// Sort strings
names := []string{"Charlie", "Alice", "Bob"}
sort.Strings(names)

// Custom sort
