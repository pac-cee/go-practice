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
			go func() {
				defer wg.Done()
				counter.Increment()
			}()
		}

		wg.Wait()
		value := counter.Value()
		if value != n {
			t.Errorf("Counter = %d; want %d", value, n)
		}
	})
}

// TestRateLimiter tests the rate limiter
func TestRateLimiter(t *testing.T) {
	t.Run("rate limiting", func(t *testing.T) {
		limiter := NewRateLimiter(100 * time.Millisecond)
		defer limiter.Stop()
