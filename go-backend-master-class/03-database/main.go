package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// User represents a user in our system
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

