package main

import (
	"reflect"
	"testing"
)

// TestDataProcessor tests the DataProcessor methods
func TestDataProcessor(t *testing.T) {
	// Test data
	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	t.Run("NewDataProcessor", func(t *testing.T) {
		processor := NewDataProcessor(data)
		if !reflect.DeepEqual(processor.data, data) {
			t.Errorf("NewDataProcessor() = %v, want %v", processor.data, data)
		}

		// Verify it's a copy
		data[0] = 100
		if processor.data[0] == 100 {
			t.Error("NewDataProcessor() did not create a copy of the data")
		}
	})

	t.Run("Sort", func(t *testing.T) {
		processor := NewDataProcessor(data)
		processor.Sort()
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		if !reflect.DeepEqual(processor.data, expected) {
			t.Errorf("Sort() = %v, want %v", processor.data, expected)
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		processor := NewDataProcessor([]int{1, 2, 3, 4, 5})
		processor.Reverse()
		expected := []int{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(processor.data, expected) {
			t.Errorf("Reverse() = %v, want %v", processor.data, expected)
		}
	})

	t.Run("Filter", func(t *testing.T) {
		processor := NewDataProcessor(data)
		evens := processor.Filter(func(x int) bool {
			return x%2 == 0
		})
		expected := []int{4, 2, 6}
		if !reflect.DeepEqual(evens, expected) {
			t.Errorf("Filter(evens) = %v, want %v", evens, expected)
		}
	})

	t.Run("Map", func(t *testing.T) {
		processor := NewDataProcessor([]int{1, 2, 3})
		doubled := processor.Map(func(x int) int {
			return x * 2
		})
		expected := []int{2, 4, 6}
		if !reflect.DeepEqual(doubled, expected) {
			t.Errorf("Map(double) = %v, want %v", doubled, expected)
		}
	})

	t.Run("RemoveDuplicates", func(t *testing.T) {
		processor := NewDataProcessor(data)
		processor.RemoveDuplicates()
		expected := []int{3, 1, 4, 5, 9, 2, 6}
		if !reflect.DeepEqual(processor.data, expected) {
			t.Errorf("RemoveDuplicates() = %v, want %v", processor.data, expected)
		}
	})

	t.Run("SlidingWindow", func(t *testing.T) {
		processor := NewDataProcessor([]int{1, 2, 3, 4, 5})
		windows := processor.SlidingWindow(3)
		expected := [][]int{
			{1, 2, 3},
			{2, 3, 4},
			{3, 4, 5},
		}
		if !reflect.DeepEqual(windows, expected) {
			t.Errorf("SlidingWindow(3) = %v, want %v", windows, expected)
		}
	})

	t.Run("Chunk", func(t *testing.T) {
		processor := NewDataProcessor([]int{1, 2, 3, 4, 5})
		chunks := processor.Chunk(2)
		expected := [][]int{
			{1, 2},
			{3, 4},
			{5},
		}
