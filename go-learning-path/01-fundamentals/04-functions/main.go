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
