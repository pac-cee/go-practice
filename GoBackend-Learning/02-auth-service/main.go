package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// User struct holds user credentials. In real apps, store users in a database.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// users map stores username to hashed password for demo purposes.
var users = map[string]string{} // username:hashedPassword
// jwtKey is the secret key used to sign JWT tokens. Keep this safe in real apps!
var jwtKey = []byte("supersecretkey")

// Claims struct is used for JWT token payload.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// signup handles POST /signup. It hashes the password and stores the user.
func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	// Hash the password with bcrypt (never store plain passwords!)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
