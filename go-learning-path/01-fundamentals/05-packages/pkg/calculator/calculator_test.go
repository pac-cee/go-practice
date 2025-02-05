package calculator

import (
	"math"
	"testing"
)

// TestCalculator tests basic calculator operations
func TestCalculator(t *testing.T) {
	// Create calculator with 2 decimal places precision
	calc := NewCalculator(2)

	// Test cases for basic operations
	tests := []struct {
		name     string
		a, b     float64
		op       func(float64, float64) float64
		expected float64
	}{
