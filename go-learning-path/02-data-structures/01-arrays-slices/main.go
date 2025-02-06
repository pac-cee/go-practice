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
	
