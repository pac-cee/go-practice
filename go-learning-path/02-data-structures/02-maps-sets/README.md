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

// Retrieve value
score := scores["Bob"]

// Check existence
score, exists := scores["Charlie"]
if !exists {
    fmt.Println("Charlie not found")
}

// Delete
delete(scores, "David")

// Length
size := len(scores)
```

### Iteration
```go
// Range over map
for key, value := range scores {
    fmt.Printf("%s: %d\n", key, value)
}

// Keys only
for key := range scores {
    fmt.Println(key)
}

// Values only
for _, value := range scores {
    fmt.Println(value)
}
```

### Map of Maps
```go
// Nested maps
graph := map[string]map[string]int{
    "A": {"B": 5, "C": 3},
    "B": {"A": 5, "C": 2},
    "C": {"A": 3, "B": 2},
}
```

## Sets

### Set Implementation
```go
// Set using map
type Set map[string]struct{}

// Create new set
func NewSet() Set {
    return make(Set)
}

// Add element
func (s Set) Add(item string) {
    s[item] = struct{}{}
}

// Remove element
func (s Set) Remove(item string) {
    delete(s, item)
}

// Contains element
func (s Set) Contains(item string) bool {
    _, exists := s[item]
    return exists
}
```

### Set Operations
```go
// Union
func (s Set) Union(other Set) Set {
    result := NewSet()
    for item := range s {
        result.Add(item)
    }
    for item := range other {
        result.Add(item)
    }
    return result
}

// Intersection
func (s Set) Intersection(other Set) Set {
    result := NewSet()
    for item := range s {
        if other.Contains(item) {
            result.Add(item)
        }
    }
    return result
}

// Difference
func (s Set) Difference(other Set) Set {
    result := NewSet()
    for item := range s {
        if !other.Contains(item) {
            result.Add(item)
        }
    }
    return result
}
```

## Concurrency

### Sync.Map
```go
// Thread-safe map
var cache sync.Map

// Store
cache.Store("key", value)

// Load
value, ok := cache.Load("key")

// Delete
cache.Delete("key")

// Load or Store
value, loaded := cache.LoadOrStore("key", defaultValue)
```

### Map Mutex
```go
type SafeMap struct {
