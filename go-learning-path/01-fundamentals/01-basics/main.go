// Package main demonstrates basic Go programming concepts
package main

import (
	"fmt"
	"strings"
	"time"
)

// Greeting returns a greeting message based on the time of day
func Greeting(name string) string {
	hour := time.Now().Hour()
	var timeOfDay string

	switch {
	case hour < 12:
		timeOfDay = "morning"
	case hour < 17:
		timeOfDay = "afternoon"
	default:
		timeOfDay = "evening"
	}

	// Demonstrate string manipulation
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Good %s, %s!", timeOfDay, name)
}

// PrintInfo demonstrates different ways to output information
func PrintInfo() {
	// Using Println
	fmt.Println("Welcome to Go Programming!")

	// Using Printf for formatted output
