package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserStore interface defines the methods for user storage
type UserStore interface {
	Create(user User) error
	Get(id int) (User, error)
	Update(user User) error
	Delete(id int) error
}

// InMemoryUserStore implements UserStore interface
type InMemoryUserStore struct {
	sync.RWMutex
	users map[int]User
}

// NewInMemoryUserStore creates a new in-memory user store
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[int]User),
	}
}

// Create adds a new user
