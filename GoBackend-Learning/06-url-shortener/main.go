package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

var (
	store = make(map[string]string)
	mu    sync.Mutex
)

func generateShortCode() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	json.NewDecoder(r.Body).Decode(&req)
	code := generateShortCode()
	mu.Lock()
