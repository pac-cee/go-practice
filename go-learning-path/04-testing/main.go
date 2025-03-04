// Package main demonstrates testing and debugging concepts in Go
package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Calculator represents a basic calculator with memory
type Calculator struct {
	memory float64
	mu     sync.RWMutex
}

// Add adds a number to memory and returns the result
func (c *Calculator) Add(x float64) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory += x
	return c.memory
}

// Subtract subtracts a number from memory and returns the result
func (c *Calculator) Subtract(x float64) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory -= x
	return c.memory
}

// Multiply multiplies memory by a number and returns the result
func (c *Calculator) Multiply(x float64) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory *= x
	return c.memory
}

// Divide divides memory by a number and returns the result
func (c *Calculator) Divide(x float64) (float64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if x == 0 {
		return 0, errors.New("division by zero")
	}
	c.memory /= x
	return c.memory, nil
}

// GetMemory returns the current value in memory
func (c *Calculator) GetMemory() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.memory
}

// Clear resets the memory to zero
func (c *Calculator) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory = 0
}

// StringProcessor provides string manipulation operations
type StringProcessor struct {
	cache map[string]string
	mu    sync.RWMutex
}

// NewStringProcessor creates a new StringProcessor
func NewStringProcessor() *StringProcessor {
	return &StringProcessor{
		cache: make(map[string]string),
	}
}

// Reverse reverses a string
func (sp *StringProcessor) Reverse(s string) string {
	// Check cache first
	sp.mu.RLock()
	if cached, ok := sp.cache[s]; ok {
		sp.mu.RUnlock()
		return cached
	}
	sp.mu.RUnlock()

	// Process if not in cache
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := string(runes)

	// Cache the result
	sp.mu.Lock()
	sp.cache[s] = result
	sp.mu.Unlock()

	return result
}

// IsPalindrome checks if a string is a palindrome
func (sp *StringProcessor) IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == sp.Reverse(s)
}

// Counter represents a thread-safe counter
type Counter struct {
	value int64
	mu    sync.Mutex
