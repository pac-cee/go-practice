package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type URLEntry struct {
	ID        int
	LongURL   string
	ShortCode string
	CreatedAt time.Time
	Clicks    int
}

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "urls.db")
	if err != nil {
		return err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		long_url TEXT NOT NULL,
		short_code TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		clicks INTEGER DEFAULT 0
	);`

	_, err = db.Exec(createTable)
	return err
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}

func createURL(longURL string) (*URLEntry, error) {
	// Check if URL already exists
	var entry URLEntry
	err := db.QueryRow("SELECT id, long_url, short_code, created_at, clicks FROM urls WHERE long_url = ?",
		longURL).Scan(&entry.ID, &entry.LongURL, &entry.ShortCode, &entry.CreatedAt, &entry.Clicks)
	if err == nil {
		return &entry, nil
	}

	// Generate new short code
	shortCode, err := generateShortCode()
	if err != nil {
		return nil, err
	}

	// Insert new URL
	result, err := db.Exec("INSERT INTO urls (long_url, short_code) VALUES (?, ?)",
		longURL, shortCode)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &URLEntry{
		ID:        int(id),
		LongURL:   longURL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		Clicks:    0,
	}, nil
}

func getURL(shortCode string) (*URLEntry, error) {
	var entry URLEntry
	err := db.QueryRow("SELECT id, long_url, short_code, created_at, clicks FROM urls WHERE short_code = ?",
		shortCode).Scan(&entry.ID, &entry.LongURL, &entry.ShortCode, &entry.CreatedAt, &entry.Clicks)
	if err != nil {
		return nil, err
	}
	return &entry, nil
