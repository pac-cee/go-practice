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
	s.RLock()
	defer s.RUnlock()
	_, exists := s.data[item]
	return exists
}

// Size returns the number of items in the set
func (s *Set) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.data)
}

// Items returns all items in the set
func (s *Set) Items() []string {
	s.RLock()
	defer s.RUnlock()
	items := make([]string, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}

// Union returns a new set containing all items from both sets
func (s *Set) Union(other *Set) *Set {
	result := NewSet()
	
	s.RLock()
	for item := range s.data {
		result.Add(item)
	}
	s.RUnlock()

	other.RLock()
	for item := range other.data {
		result.Add(item)
	}
