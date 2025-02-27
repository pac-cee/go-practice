// Package main demonstrates concurrency patterns in Go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SafeCounter is a thread-safe counter using channels
type SafeCounter struct {
	value chan int
	read  chan chan int
}

// NewSafeCounter creates a new thread-safe counter
func NewSafeCounter() *SafeCounter {
	c := &SafeCounter{
		value: make(chan int),
		read:  make(chan chan int),
	}
	go c.run()
	return c
}

func (c *SafeCounter) run() {
	value := 0
	for {
		select {
		case <-c.value:
			value++
		case reader := <-c.read:
			reader <- value
		}
	}
}

// Increment increases the counter by one
func (c *SafeCounter) Increment() {
	c.value <- 1
}

// Value returns the current value of the counter
func (c *SafeCounter) Value() int {
	reader := make(chan int)
	c.read <- reader
	return <-reader
}

// RateLimiter implements a simple rate limiter
type RateLimiter struct {
	ticker *time.Ticker
	stop   chan struct{}
}

