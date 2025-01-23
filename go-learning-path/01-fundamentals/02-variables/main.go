// Package main demonstrates variable usage and type manipulation in Go
package main

import (
	"fmt"
	"strconv"
)

// UserScore represents a user's game score
type UserScore struct {
	Username string
	Score    int
	Rank     float64
	Active   bool
}

// Define constants for game levels
const (
	BeginnerLevel = iota
	IntermediateLevel
	AdvancedLevel
	ExpertLevel
)

// calculateRank determines a user's rank based on their score
func calculateRank(score int) float64 {
	return float64(score) / 100.0
}

// demonstrateVariables shows different ways to declare and use variables
func demonstrateVariables() {
	// Standard declaration
	var name string
	name = "Gopher"

	// Declaration with initial value
	var age int = 25

	// Short declaration
	score := 95

	// Multiple declaration
	var (
		isActive bool = true
		level    int  = IntermediateLevel
	)

	// Print variables
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Active: %v\n", isActive)
	fmt.Printf("Level: %d\n", level)
}

// demonstrateTypeConversion shows type conversion examples
func demonstrateTypeConversion() {
	// Numeric conversions
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Integer: %d\n", i)
	fmt.Printf("Float: %f\n", f)
	fmt.Printf("Unsigned Integer: %d\n", u)

	// String conversions
	str := strconv.Itoa(i)
	fmt.Printf("String: %s\n", str)

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error converting string to integer: %v\n", err)
	} else {
		fmt.Printf("Back to number: %d\n", num)
	}
