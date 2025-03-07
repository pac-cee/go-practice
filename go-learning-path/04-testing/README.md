# Testing and Debugging in Go 🧪

## 📖 Table of Contents
1. [Unit Testing](#unit-testing)
2. [Table-Driven Tests](#table-driven-tests)
3. [Benchmarking](#benchmarking)
4. [Profiling](#profiling)
5. [Race Detection](#race-detection)
6. [Debugging](#debugging)
7. [Best Practices](#best-practices)
8. [Exercises](#exercises)

## Unit Testing

### Test File Structure
```go
package mypackage_test  // Use separate package for black-box testing

import (
