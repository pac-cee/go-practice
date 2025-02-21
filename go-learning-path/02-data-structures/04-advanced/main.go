// Package main demonstrates advanced data structures in Go
package main

import (
	"container/list"
	"errors"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
)

// MinHeap implements a min heap data structure
type MinHeap struct {
	items []int
	size  int
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: make([]int, 0),
	}
}

// Insert adds an item to the heap
func (h *MinHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.size++
	h.heapifyUp(h.size - 1)
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() (int, error) {
	if h.size == 0 {
		return 0, errors.New("heap is empty")
	}

	min := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.items = h.items[:h.size]
	if h.size > 0 {
		h.heapifyDown(0)
	}

	return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[parent] > h.items[index] {
			h.items[parent], h.items[index] = h.items[index], h.items[parent]
			index = parent
		} else {
