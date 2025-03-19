package main

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/blog_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Create migration manager and run migrations
	migrationManager := NewMigrationManager(db)
	if err := migrationManager.Migrate(); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return db, func() {
		// Clean up database
		db.Exec("DROP TABLE comments")
		db.Exec("DROP TABLE posts")
		db.Exec("DROP TABLE schema_migrations")
		db.Close()
	}
}

func TestBlogStore(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	store := NewBlogStore(db)
	ctx := context.Background()

	t.Run("create and get post", func(t *testing.T) {
		post := &Post{
			Title:   "Test Post",
			Content: "Test Content",
		}

		// Create post
		err := store.CreatePost(ctx, post)
		if err != nil {
			t.Fatalf("Failed to create post: %v", err)
		}
		if post.ID == 0 {
