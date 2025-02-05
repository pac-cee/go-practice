// Package calculator provides basic and advanced mathematical operations.
// It demonstrates proper package organization, documentation, and testing.
package calculator

import (
	"errors"
	"math"
)

// Common errors that can be returned by the calculator package
var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNegativeNumber = errors.New("negative number not allowed")
)

// Operation represents a mathematical operation
type Operation func(float64, float64) (float64, error)

// Calculator provides mathematical operations with optional logging
