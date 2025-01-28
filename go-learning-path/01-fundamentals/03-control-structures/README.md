# Control Structures in Go 🔄

## 📖 Table of Contents
1. [If Statements](#if-statements)
2. [For Loops](#for-loops)
3. [Switch Statements](#switch-statements)
4. [Defer](#defer)
5. [Panic and Recover](#panic-and-recover)
6. [Best Practices](#best-practices)
7. [Exercises](#exercises)

## If Statements

### Basic If Statement
```go
if condition {
    // code
}
```

### If-Else Statement
```go
if condition {
    // code
} else {
    // code
}
```

### If with Initialization
```go
if value := getValue(); value > 10 {
    // code
}
```

### Multiple Conditions
```go
if condition1 && condition2 {
    // both true
} else if condition1 || condition2 {
    // at least one true
} else {
    // none true
}
```

## For Loops

### Basic For Loop
```go
for i := 0; i < 10; i++ {
    // code
}
```

### While-style Loop
```go
for condition {
    // code
}
```

### Infinite Loop
```go
for {
    // code
    if shouldBreak {
        break
    }
}
```

### Range Loop
```go
// Slice/Array
for index, value := range slice {
    // code
}

// Map
for key, value := range map {
    // code
}

// String
for index, char := range string {
    // code
}

// Channel
for value := range channel {
    // code
}
```
