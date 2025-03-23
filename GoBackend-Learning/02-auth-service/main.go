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

