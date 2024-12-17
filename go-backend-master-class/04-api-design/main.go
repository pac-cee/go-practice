package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Product represents a product in our system
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ProductStore is an in-memory store for products
type ProductStore struct {
	sync.RWMutex
	products map[string]Product
}

// NewProductStore creates a new product store
func NewProductStore() *ProductStore {
	return &ProductStore{
		products: make(map[string]Product),
