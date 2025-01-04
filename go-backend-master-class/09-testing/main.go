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
