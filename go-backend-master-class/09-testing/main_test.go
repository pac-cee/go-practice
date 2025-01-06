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
}

func (s *MockUserStore) Create(user User) error {
	s.users[user.ID] = user
	return nil
}

func (s *MockUserStore) Get(id int) (User, error) {
	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *MockUserStore) Update(user User) error {
	s.users[user.ID] = user
	return nil
}

func (s *MockUserStore) Delete(id int) error {
	delete(s.users, id)
	return nil
}

// TestUserService_CreateUser tests the CreateUser method
func TestUserService_CreateUser(t *testing.T) {
	store := NewMockUserStore()
	service := NewUserService(store)

	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user",
			user: User{
