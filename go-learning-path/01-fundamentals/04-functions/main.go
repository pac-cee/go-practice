// Package main demonstrates functions and methods in Go
package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

// Shape interface defines common methods for shapes
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Circle implements Shape
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.radius)
}

// Rectangle implements Shape
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f)", r.width, r.height)
}

// Logger interface for different logging implementations
type Logger interface {
	Log(message string)
	Error(message string)
}

// ConsoleLogger implements Logger
type ConsoleLogger struct {
	prefix string
}

func (l ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] LOG: %s\n", l.prefix, message)
}

func (l ConsoleLogger) Error(message string) {
	fmt.Printf("[%s] ERROR: %s\n", l.prefix, message)
}

// Calculator demonstrates function types and closures
type Calculator struct {
	logger Logger
}

// Operation represents a mathematical operation
type Operation func(a, b float64) float64

// Basic operations
func add(a, b float64) float64      { return a + b }
func subtract(a, b float64) float64 { return a - b }
func multiply(a, b float64) float64 { return a * b }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Calculate performs the operation and logs the result
func (c Calculator) Calculate(a, b float64, op Operation) float64 {
	result := op(a, b)
	c.logger.Log(fmt.Sprintf("Calculated result: %.2f", result))
