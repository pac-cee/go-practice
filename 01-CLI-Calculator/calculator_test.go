package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"5 + 3", 8, false},
		{"10 - 5", 5, false},
		{"4 * 3", 12, false},
		{"15 / 3", 5, false},
		{"10 / 0", 0, true},     // Division by zero
		{"abc + 3", 0, true},    // Invalid number
		{"5 & 3", 0, true},      // Invalid operator
