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
