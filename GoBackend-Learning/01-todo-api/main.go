package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Todo represents a single to-do item.
// Struct tags (e.g. `json:"id"`) tell Go how to encode/decode fields to/from JSON.
type Todo struct {
	ID   int    `json:"id"`   // Unique identifier
	Task string `json:"task"` // Task description
	Done bool   `json:"done"` // Completion status
}

// In-memory slice to store to-do items. In real apps, you'd use a database.
var todos = []Todo{
