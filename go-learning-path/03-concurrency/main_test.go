package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestSafeCounter tests the thread-safe counter
func TestSafeCounter(t *testing.T) {
	t.Run("concurrent increments", func(t *testing.T) {
		counter := NewSafeCounter()
		var wg sync.WaitGroup
		n := 100

		// Launch n goroutines to increment
		for i := 0; i < n; i++ {
			wg.Add(1)
