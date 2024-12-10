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
