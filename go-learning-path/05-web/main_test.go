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
