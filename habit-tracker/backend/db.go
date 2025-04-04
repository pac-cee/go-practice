package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "habit_tracker.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	// Create tables if they don't exist
	_, err = db.Exec(`
