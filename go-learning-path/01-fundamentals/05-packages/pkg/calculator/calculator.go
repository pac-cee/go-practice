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
type Calculator struct {
	// EnableLogging determines if operations should be logged
	EnableLogging bool
	
	// precision specifies decimal places for rounding
	precision int
}

// NewCalculator creates a new Calculator with the specified precision
func NewCalculator(precision int) *Calculator {
	return &Calculator{
		precision: precision,
	}
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return c.round(a + b)
}

// Subtract returns the difference between two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return c.round(a - b)
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return c.round(a * b)
}

// Divide returns the quotient of two numbers
// Returns an error if dividing by zero
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return c.round(a / b), nil
}
