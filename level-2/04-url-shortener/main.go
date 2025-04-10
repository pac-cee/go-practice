package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type URLEntry struct {
	ID        int
	LongURL   string
