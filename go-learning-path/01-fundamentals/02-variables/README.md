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
