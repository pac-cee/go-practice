// Package main demonstrates maps and sets in Go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Set represents a thread-safe set data structure
type Set struct {
	sync.RWMutex
	data map[string]struct{}
}

// NewSet creates a new empty set
func NewSet() *Set {
	return &Set{
		data: make(map[string]struct{}),
	}
}

// Add adds an item to the set
func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.data[item] = struct{}{}
}

// Remove removes an item from the set
func (s *Set) Remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, item)
}

// Contains checks if an item exists in the set
func (s *Set) Contains(item string) bool {
