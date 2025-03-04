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
		if got := calc.Subtract(3); got != 2 {
			t.Errorf("Subtract(3) = %.2f; want 2", got)
		}

		// Test Multiply
		if got := calc.Multiply(4); got != 8 {
			t.Errorf("Multiply(4) = %.2f; want 8", got)
		}

		// Test Divide
		if got, err := calc.Divide(2); err != nil || got != 4 {
			t.Errorf("Divide(2) = %.2f, %v; want 4, nil", got, err)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		calc := &Calculator{}
		calc.Add(10)
		if _, err := calc.Divide(0); err == nil {
			t.Error("Divide(0) should return error")
		}
	})

	t.Run("concurrent operations", func(t *testing.T) {
		calc := &Calculator{}
		var wg sync.WaitGroup
		n := 100

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				calc.Add(1)
			}()
		}
		wg.Wait()

		if got := calc.GetMemory(); got != float64(n) {
			t.Errorf("After %d concurrent adds, memory = %.2f; want %d", n, got, n)
		}
	})
}

// TestStringProcessor tests the StringProcessor type
func TestStringProcessor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"simple string", "hello", "olleh"},
		{"unicode string", "世界", "界世"},
	}

	sp := NewStringProcessor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sp.Reverse(tt.input); got != tt.want {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}

	t.Run("palindrome check", func(t *testing.T) {
		palindromes := []struct {
			input string
			want  bool
		}{
			{"", true},
			{"a", true},
			{"race a car", false},
			{"A man a plan a canal Panama", true},
			{"Was it a car or a cat I saw", true},
		}
