package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Set DATABASE_URL env var to use Postgres, e.g.
// postgres://user:password@host:port/dbname?sslmode=disable

var db *gorm.DB

func getTodos(w http.ResponseWriter, r *http.Request) {
