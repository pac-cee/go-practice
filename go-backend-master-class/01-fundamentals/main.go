package main

import (
	"fmt"
	"sync"
	"time"
)

// User represents a basic user structure
type User struct {
	ID       int
	Username string
	Email    string
}

// UserInterface defines behavior for user operations
type UserInterface interface {
	GetEmail() string
	UpdateEmail(newEmail string)
}

// GetEmail returns user's email
func (u *User) GetEmail() string {
	return u.Email
}

// UpdateEmail updates user's email
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// SafeCounter is a thread-safe counter using mutex
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Increment increases the counter by 1
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

// GetCount returns the current count
func (sc *SafeCounter) GetCount() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

// processUser demonstrates error handling
func processUser(u *User) error {
	if u == nil {
		return fmt.Errorf("user cannot be nil")
	}
	if u.Email == "" {
