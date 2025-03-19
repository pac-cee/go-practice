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
		post.Content,
		post.CreatedAt,
		post.UpdatedAt,
	).Scan(&post.ID)
}

// GetPost retrieves a post by ID
func (s *BlogStore) GetPost(ctx context.Context, id int) (*Post, error) {
	post := &Post{}
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM posts
		WHERE id = $1`

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found: %d", id)
	}
	return post, err
}

// UpdatePost updates an existing post
func (s *BlogStore) UpdatePost(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4`

	post.UpdatedAt = time.Now()
	result, err := s.db.ExecContext(
		ctx, query,
		post.Title,
		post.Content,
		post.UpdatedAt,
		post.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("post not found: %d", post.ID)
	}
	return nil
}

// DeletePost deletes a post and its comments
func (s *BlogStore) DeletePost(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Delete comments first
	_, err = tx.ExecContext(ctx, "DELETE FROM comments WHERE post_id = $1", id)
	if err != nil {
		return err
	}

	// Delete the post
	result, err := tx.ExecContext(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("post not found: %d", id)
	}

	return tx.Commit()
}

// AddComment adds a comment to a post
func (s *BlogStore) AddComment(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments (post_id, content, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	comment.CreatedAt = time.Now()

	return s.db.QueryRowContext(
		ctx, query,
		comment.PostID,
		comment.Content,
		comment.CreatedAt,
	).Scan(&comment.ID)
}

// GetPostComments retrieves all comments for a post
func (s *BlogStore) GetPostComments(ctx context.Context, postID int) ([]Comment, error) {
	query := `
		SELECT id, post_id, content, created_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC`

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
