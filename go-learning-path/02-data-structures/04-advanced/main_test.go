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
		expected := []int{1, 3, 4, 5, 7}
		for _, want := range expected {
			got, err := heap.ExtractMin()
			if err != nil {
				t.Errorf("ExtractMin() error = %v", err)
			}
			if got != want {
				t.Errorf("ExtractMin() = %v, want %v", got, want)
			}
		}

		// Test empty heap
		_, err := heap.ExtractMin()
		if err == nil {
			t.Error("ExtractMin() on empty heap should return error")
		}
	})
}

