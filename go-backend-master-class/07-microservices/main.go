package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Service represents a microservice
type Service struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Status   string `json:"status"`
	LastPing time.Time `json:"last_ping"`
}

// ServiceRegistry keeps track of available services
type ServiceRegistry struct {
	Services map[string]Service
}

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewServiceRegistry creates a new service registry
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		Services: make(map[string]Service),
