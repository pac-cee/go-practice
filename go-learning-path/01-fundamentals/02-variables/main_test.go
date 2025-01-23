package main

import (
	"testing"
)

// TestCalculateRank tests the rank calculation function
func TestCalculateRank(t *testing.T) {
	tests := []struct {
		name     string
		score    int
		expected float64
	}{
		{
			name:     "perfect score",
			score:    100,
			expected: 1.0,
		},
		{
