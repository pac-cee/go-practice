// Package main demonstrates web development concepts in Go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Todo represents a todo item
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoStore manages todo items
type TodoStore struct {
	sync.RWMutex
	todos map[int]*Todo
	seq   int
}

// NewTodoStore creates a new todo store
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos: make(map[int]*Todo),
	}
}

// Create adds a new todo
func (s *TodoStore) Create(title string) *Todo {
	s.Lock()
	defer s.Unlock()

	s.seq++
	todo := &Todo{
		ID:        s.seq,
		Title:     title,
		CreatedAt: time.Now(),
	}
	s.todos[todo.ID] = todo
	return todo
}

// Get retrieves a todo by ID
func (s *TodoStore) Get(id int) (*Todo, bool) {
	s.RLock()
