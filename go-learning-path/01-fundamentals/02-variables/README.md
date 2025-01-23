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
