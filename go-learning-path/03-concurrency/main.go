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

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit time.Duration) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(limit),
		stop:   make(chan struct{}),
	}
}

// Allow returns a channel that receives when action is allowed
func (r *RateLimiter) Allow() <-chan time.Time {
	return r.ticker.C
}

// Stop stops the rate limiter
func (r *RateLimiter) Stop() {
	r.ticker.Stop()
	close(r.stop)
}

// Worker represents a worker pool worker
type Worker struct {
	id      int
	jobs    <-chan int
	results chan<- int
}

// NewWorker creates a new worker
func NewWorker(id int, jobs <-chan int, results chan<- int) *Worker {
	return &Worker{
		id:      id,
		jobs:    jobs,
		results: results,
	}
}

// Start starts the worker
func (w *Worker) Start() {
	go func() {
		for j := range w.jobs {
			fmt.Printf("worker %d processing job %d\n", w.id, j)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			w.results <- j * 2
		}
	}()
}

// Pipeline demonstrates a simple pipeline pattern
type Pipeline struct {
	done <-chan struct{}
}

// NewPipeline creates a new pipeline
func NewPipeline(done <-chan struct{}) *Pipeline {
	return &Pipeline{done: done}
}

// Generator generates integers from 1 to n
func (p *Pipeline) Generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			select {
			case <-p.done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

// Square squares numbers from in channel
func (p *Pipeline) Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-p.done:
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

// Filter filters out odd numbers
func (p *Pipeline) Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				select {
				case <-p.done:
					return
				case out <- n:
				}
			}
		}
	}()
	return out
}

// FanOut demonstrates the fan-out pattern
func FanOut(in <-chan int, n int) []<-chan int {
	channels := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = processChannel(in)
	}
	return channels
}
