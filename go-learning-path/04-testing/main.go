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
