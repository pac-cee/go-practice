package main

import (
	"reflect"
	"testing"
)

// TestMinHeap tests heap operations
func TestMinHeap(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		heap := NewMinHeap()
		values := []int{5, 3, 7, 1, 4}

		// Test Insert
		for _, v := range values {
			heap.Insert(v)
		}

		// Test ExtractMin
