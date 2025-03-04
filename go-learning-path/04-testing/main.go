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
