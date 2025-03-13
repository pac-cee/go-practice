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
	})

	t.Run("list todos", func(t *testing.T) {
		store.Create("Another todo")
		todos := store.List()
		if len(todos) != 2 {
			t.Errorf("expected 2 todos, got %d", len(todos))
		}
	})

	t.Run("update todo", func(t *testing.T) {
		success := store.Update(1, true)
		if !success {
			t.Error("update should succeed")
		}

		todo, _ := store.Get(1)
		if !todo.Completed {
			t.Error("todo should be completed")
		}
	})

	t.Run("delete todo", func(t *testing.T) {
		success := store.Delete(1)
		if !success {
			t.Error("delete should succeed")
		}

		_, exists := store.Get(1)
		if exists {
			t.Error("todo should not exist after deletion")
		}
	})
}

func TestServer(t *testing.T) {
	server := NewServer()
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	t.Run("get todos", func(t *testing.T) {
		resp, err := client.Get(ts.URL + "/api/todos")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK, got %v", resp.Status)
		}

		var todos []*Todo
		if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("create todo", func(t *testing.T) {
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "Test todo",
		}

		body, _ := json.Marshal(todo)
		resp, err := client.Post(ts.URL+"/api/todos", "application/json", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status Created, got %v", resp.Status)
		}

		var created Todo
		if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
			t.Fatal(err)
		}

		if created.Title != todo.Title {
			t.Errorf("expected title %q, got %q", todo.Title, created.Title)
		}
	})

	t.Run("update todo", func(t *testing.T) {
		// First create a todo
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "Update test",
		}
		body, _ := json.Marshal(todo)
		resp, _ := client.Post(ts.URL+"/api/todos", "application/json", bytes.NewReader(body))
		var created Todo
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()

		// Now update it
		update := struct {
			Completed bool `json:"completed"`
		}{
			Completed: true,
		}
		body, _ = json.Marshal(update)
		req, _ := http.NewRequest(http.MethodPut, ts.URL+"/api/todos/"+string(created.ID), bytes.NewReader(body))
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("expected status No Content, got %v", resp.Status)
		}
	})

	t.Run("delete todo", func(t *testing.T) {
		// First create a todo
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "Delete test",
		}
		body, _ := json.Marshal(todo)
		resp, _ := client.Post(ts.URL+"/api/todos", "application/json", bytes.NewReader(body))
		var created Todo
		json.NewDecoder(resp.Body).Decode(&created)
		resp.Body.Close()
