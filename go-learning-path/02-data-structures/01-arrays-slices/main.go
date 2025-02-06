// Package main demonstrates array and slice operations in Go
package main

import (
	"fmt"
	"sort"
)

// DataProcessor demonstrates various array and slice operations
type DataProcessor struct {
	data []int
}

// NewDataProcessor creates a new DataProcessor with initial data
func NewDataProcessor(data []int) *DataProcessor {
	// Create a copy of the input data
	copiedData := make([]int, len(data))
	copy(copiedData, data)
	
	return &DataProcessor{
		data: copiedData,
	}
}

// Sort sorts the data in ascending order
func (dp *DataProcessor) Sort() {
	sort.Ints(dp.data)
}

// Reverse reverses the order of elements
func (dp *DataProcessor) Reverse() {
	for i := 0; i < len(dp.data)/2; i++ {
		j := len(dp.data) - 1 - i
		dp.data[i], dp.data[j] = dp.data[j], dp.data[i]
	}
}

// Filter returns elements that satisfy the predicate
