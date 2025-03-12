package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestTodoStore(t *testing.T) {
	store := NewTodoStore()

	t.Run("create todo", func(t *testing.T) {
		todo := store.Create("Test todo")
		if todo.Title != "Test todo" {
			t.Errorf("expected title 'Test todo', got %q", todo.Title)
		}
		if todo.ID != 1 {
			t.Errorf("expected ID 1, got %d", todo.ID)
		}
		if todo.Completed {
			t.Error("new todo should not be completed")
		}
	})

	t.Run("get todo", func(t *testing.T) {
		todo, exists := store.Get(1)
		if !exists {
			t.Error("todo should exist")
		}
		if todo.Title != "Test todo" {
			t.Errorf("expected title 'Test todo', got %q", todo.Title)
		}
