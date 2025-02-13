# Maps and Sets in Go 🗺️

## 📖 Table of Contents
1. [Maps](#maps)
2. [Sets](#sets)
3. [Operations](#operations)
4. [Concurrency](#concurrency)
5. [Best Practices](#best-practices)
6. [Exercises](#exercises)

## Maps

### Map Declaration
```go
// Empty map
var scores map[string]int

// Map with make
users := make(map[string]User)

// Map literal
capitals := map[string]string{
    "USA": "Washington",
    "UK":  "London",
    "FR":  "Paris",
}
```

### Map Properties
- Key-value pairs
- Unordered collection
- Reference type
- Fast lookups (O(1) average)
- Dynamic size

### Key Types
Valid key types:
- Comparable types
- Numbers
- Strings
- Pointers
- Interfaces
- Structs
- Arrays

Invalid key types:
- Slices
- Maps
- Functions

## Operations

### Basic Operations
```go
// Insert or update
scores["Alice"] = 100

