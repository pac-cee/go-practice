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
