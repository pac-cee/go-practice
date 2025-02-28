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

		start := time.Now()
		count := 0
		timeout := time.After(550 * time.Millisecond)

		// Should allow ~5 actions in 500ms
		for {
			select {
			case <-timeout:
				if count < 4 || count > 6 {
					t.Errorf("Got %d actions; want 5 ±1", count)
				}
				return
			case <-limiter.Allow():
				count++
			}
		}
	})
}
