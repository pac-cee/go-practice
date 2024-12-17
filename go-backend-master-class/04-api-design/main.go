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
	}
}

// API represents our REST API
type API struct {
	store *ProductStore
}

// NewAPI creates a new API instance
func NewAPI(store *ProductStore) *API {
	return &API{store: store}
}

// respondWithJSON writes JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// CreateProduct handles product creation
func (api *API) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

	product.CreatedAt = time.Now()
	product.ID = time.Now().Format("20060102150405")

	api.store.Lock()
	api.store.products[product.ID] = product
	api.store.Unlock()

	respondWithJSON(w, http.StatusCreated, APIResponse{
		Status:  "success",
		Message: "Product created successfully",
		Data:    product,
	})
}

// GetProduct handles getting a single product
func (api *API) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	api.store.RLock()
	product, exists := api.store.products[id]
	api.store.RUnlock()

	if !exists {
		respondWithJSON(w, http.StatusNotFound, APIResponse{
			Status:  "error",
			Message: "Product not found",
		})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Status: "success",
		Data:   product,
	})
}

// ListProducts handles listing all products
func (api *API) ListProducts(w http.ResponseWriter, r *http.Request) {
	api.store.RLock()
	products := make([]Product, 0, len(api.store.products))
	for _, product := range api.store.products {
		products = append(products, product)
	}
	api.store.RUnlock()

	respondWithJSON(w, http.StatusOK, APIResponse{
		Status: "success",
		Data:   products,
	})
}

// UpdateProduct handles updating a product
func (api *API) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

	api.store.Lock()
	if _, exists := api.store.products[id]; !exists {
		api.store.Unlock()
		respondWithJSON(w, http.StatusNotFound, APIResponse{
			Status:  "error",
			Message: "Product not found",
		})
		return
	}

	// Preserve created time and ID
	product.ID = id
	product.CreatedAt = api.store.products[id].CreatedAt
	api.store.products[id] = product
	api.store.Unlock()

