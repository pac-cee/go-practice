package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

// User represents the user model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
