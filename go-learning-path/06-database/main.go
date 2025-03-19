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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comment represents a comment on a post
type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// BlogStore handles database operations for the blog
type BlogStore struct {
	db *sql.DB
}

// NewBlogStore creates a new BlogStore
func NewBlogStore(db *sql.DB) *BlogStore {
	return &BlogStore{db: db}
}

// CreatePost creates a new blog post
func (s *BlogStore) CreatePost(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	return s.db.QueryRowContext(
		ctx, query,
		post.Title,
