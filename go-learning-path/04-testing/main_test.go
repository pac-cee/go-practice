package main

import (
	"fmt"
	"sync"
	"testing"
)

// TestCalculator tests the Calculator type
func TestCalculator(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		calc := &Calculator{}

		// Test Add
		if got := calc.Add(5); got != 5 {
			t.Errorf("Add(5) = %.2f; want 5", got)
		}

		// Test Subtract
