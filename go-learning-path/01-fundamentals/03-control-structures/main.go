// Package main demonstrates control structures in Go
package main

import (
	"fmt"
	"time"
)

// UserRole represents user access levels
type UserRole int

const (
	Guest UserRole = iota
	User
	Moderator
	Admin
)

// User represents a system user
