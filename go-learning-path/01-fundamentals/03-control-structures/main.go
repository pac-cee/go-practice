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
type User struct {
	ID       int
	Name     string
	Role     UserRole
	Active   bool
	JoinDate time.Time
}

// checkUserAccess demonstrates if statements and switch
func checkUserAccess(u User, resource string) {
	// If with initialization
	if lastLogin := time.Since(u.JoinDate); lastLogin.Hours() < 24 {
		fmt.Printf("New user detected: %s (joined %v ago)\n", 
			u.Name, lastLogin.Round(time.Minute))
	}

	// Switch on user role
	switch u.Role {
	case Admin:
