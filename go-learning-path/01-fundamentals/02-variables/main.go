// Package main demonstrates variable usage and type manipulation in Go
package main

import (
	"fmt"
	"strconv"
)

// UserScore represents a user's game score
type UserScore struct {
	Username string
	Score    int
	Rank     float64
	Active   bool
}

// Define constants for game levels
const (
	BeginnerLevel = iota
