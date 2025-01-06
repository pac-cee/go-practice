package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockUserStore implements UserStore interface for testing
type MockUserStore struct {
	users map[int]User
}

func NewMockUserStore() *MockUserStore {
	return &MockUserStore{
		users: make(map[int]User),
	}
