# Variables, Types, and Constants in Go 🔄

## 📖 Table of Contents
1. [Basic Types](#basic-types)
2. [Variable Declaration](#variable-declaration)
3. [Type Conversion](#type-conversion)
4. [Constants](#constants)
5. [Type Inference](#type-inference)
6. [Zero Values](#zero-values)
7. [Composite Types](#composite-types)
8. [Exercises](#exercises)

## Basic Types

Go has several basic types:

```go
// Numeric Types
int8, int16, int32, int64
uint8, uint16, uint32, uint64
int  // platform dependent, 32 or 64 bit
uint // platform dependent, 32 or 64 bit
float32, float64
complex64, complex128

// Text Types
string
rune   // alias for int32, represents a Unicode code point
byte   // alias for uint8

// Boolean Type
bool

// Error Type
error
```

### Size and Range Examples
```go
// Integer ranges
int8:   -128 to 127
int16:  -32768 to 32767
int32:  -2147483648 to 2147483647
int64:  -9223372036854775808 to 9223372036854775807

// Unsigned integer ranges
uint8:  0 to 255
uint16: 0 to 65535
uint32: 0 to 4294967295
uint64: 0 to 18446744073709551615
```

## Variable Declaration

There are several ways to declare variables in Go:

### 1. Standard Declaration
```go
var name string
var age int
var isActive bool

// Multiple declarations
var (
    firstName string
    lastName  string
    score     int
)
```

### 2. Declaration with Initial Value
```go
var name string = "Gopher"
var age int = 25
var isActive = true  // Type inference
```

### 3. Short Declaration (Inside Functions)
```go
func example() {
    name := "Gopher"    // string
    age := 25          // int
    score := 95.5      // float64
}
```

### 4. Multiple Assignments
```go
var x, y int = 10, 20
name, age := "Gopher", 25
```

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String conversions
str := strconv.Itoa(i)      // int to string
num, err := strconv.Atoi(str) // string to int

// Float conversions
s := fmt.Sprintf("%.2f", f) // float to string
f, err = strconv.ParseFloat(s, 64) // string to float
```

## Constants

Constants are declared using the `const` keyword:

```go
// Single constant
const Pi = 3.14159

// Multiple constants
const (
    StatusOK    = 200
    StatusError = 500
)

// iota for enumerated constants
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday        // 5
    Saturday      // 6
