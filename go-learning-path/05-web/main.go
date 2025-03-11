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
	defer s.RUnlock()
	todo, ok := s.todos[id]
	return todo, ok
}

// List returns all todos
func (s *TodoStore) List() []*Todo {
	s.RLock()
	defer s.RUnlock()
	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

// Update modifies a todo
func (s *TodoStore) Update(id int, completed bool) bool {
	s.Lock()
	defer s.Unlock()
	if todo, ok := s.todos[id]; ok {
		todo.Completed = completed
		return true
	}
	return false
}

// Delete removes a todo
func (s *TodoStore) Delete(id int) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.todos[id]; ok {
		delete(s.todos, id)
		return true
	}
	return false
}

// Server represents the web server
type Server struct {
	router    *mux.Router
	todos     *TodoStore
	templates *template.Template
	upgrader  websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan Message
}

// Message represents a WebSocket message
type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	User    string `json:"user"`
}

// NewServer creates a new web server
func NewServer() *Server {
	s := &Server{
		router:  mux.NewRouter(),
		todos:   NewTodoStore(),
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // In production, check origin
			},
		},
	}

	// Parse templates
	var err error
	s.templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
