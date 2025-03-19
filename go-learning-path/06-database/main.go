// Package main demonstrates database integration concepts in Go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// Post represents a blog post
type Post struct {
	ID        int       `json:"id"`
