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
