package main

import (
	"strings"
	"testing"
)

// TestGreeting tests the Greeting function
func TestGreeting(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{
			name:     "normal name",
			input:    "Alice",
			contains: "Alice",
		},
		{
			name:     "empty string",
			input:    "",
			contains: "World",
		},
		{
			name:     "with spaces",
			input:    "  Bob  ",
			contains: "Bob",
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greeting(tt.input)
			if !strings.Contains(result, tt.contains) {
				t.Errorf("Greeting(%q) = %q, want it to contain %q", 
