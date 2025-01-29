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
