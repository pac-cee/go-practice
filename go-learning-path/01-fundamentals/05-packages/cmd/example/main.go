// Package main demonstrates the usage of the calculator package
package main

import (
	"fmt"
	"log"

	"go-learning-path/01-fundamentals/05-packages/pkg/calculator"
)

func main() {
	// Create a new calculator with 2 decimal places precision
	calc := calculator.NewCalculator(2)

	// Enable logging
	calc.EnableLogging = true

	// Basic operations
	a, b := 10.0, 5.0
