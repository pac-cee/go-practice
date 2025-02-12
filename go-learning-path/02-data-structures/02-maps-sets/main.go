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
	other.RUnlock()

	return result
}

// Intersection returns a new set containing items present in both sets
func (s *Set) Intersection(other *Set) *Set {
	result := NewSet()
	
	s.RLock()
	defer s.RUnlock()
	
	other.RLock()
	defer other.RUnlock()

	// Iterate over the smaller set for efficiency
	if len(s.data) < len(other.data) {
		for item := range s.data {
			if _, exists := other.data[item]; exists {
				result.Add(item)
			}
		}
	} else {
		for item := range other.data {
			if _, exists := s.data[item]; exists {
				result.Add(item)
			}
		}
	}

	return result
}

// Difference returns a new set containing items in s that are not in other
func (s *Set) Difference(other *Set) *Set {
	result := NewSet()
	
	s.RLock()
	defer s.RUnlock()
	
	other.RLock()
	defer other.RUnlock()

	for item := range s.data {
		if _, exists := other.data[item]; !exists {
			result.Add(item)
		}
	}

	return result
}

// Cache represents a thread-safe cache with expiration
type Cache struct {
	sync.RWMutex
	data    map[string]interface{}
	expires map[string]time.Time
}

// NewCache creates a new cache
func NewCache() *Cache {
	cache := &Cache{
		data:    make(map[string]interface{}),
		expires: make(map[string]time.Time),
	}
	go cache.cleanup()
	return cache
}

// Set adds an item to the cache with expiration
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
	c.expires[key] = time.Now().Add(ttl)
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	
	expiry, exists := c.expires[key]
	if !exists || time.Now().After(expiry) {
		return nil, false
	}
	
	value, exists := c.data[key]
	return value, exists
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
	delete(c.expires, key)
}

// cleanup periodically removes expired items
func (c *Cache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.Lock()
		now := time.Now()
		for key, expiry := range c.expires {
			if now.After(expiry) {
				delete(c.data, key)
				delete(c.expires, key)
			}
		}
		c.Unlock()
	}
}

// FrequencyCounter counts occurrences of words
type FrequencyCounter struct {
	sync.RWMutex
	counts map[string]int
}

// NewFrequencyCounter creates a new frequency counter
func NewFrequencyCounter() *FrequencyCounter {
	return &FrequencyCounter{
		counts: make(map[string]int),
	}
}

// Add increments the count for a word
func (fc *FrequencyCounter) Add(word string) {
	fc.Lock()
	defer fc.Unlock()
	fc.counts[word]++
