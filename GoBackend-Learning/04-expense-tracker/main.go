package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Expense struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

var db *sql.DB
