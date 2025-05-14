package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Todo{})
	return db
}

func TestRegisterLoginAndTodoCRUD(t *testing.T) {
	db := setupTestDB()
	// Register
	r := httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"testuser","password":"testpass"}`))
	w := httptest.NewRecorder()
	register(db)(w, r)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
	// Login
	r = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"testpass"}`))
	w = httptest.NewRecorder()
	login(db)(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp AuthResponse
	_ = json.NewDecoder(w.Body).Decode(&resp)
	if resp.Token == "" {
		t.Fatal("expected token")
