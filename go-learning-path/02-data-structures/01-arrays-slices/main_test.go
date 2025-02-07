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
