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
	}
}

// RegisterService registers a new service
func (sr *ServiceRegistry) RegisterService(w http.ResponseWriter, r *http.Request) {
	var service Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

	service.Status = "healthy"
	service.LastPing = time.Now()
	sr.Services[service.Name] = service

	respondWithJSON(w, http.StatusCreated, Response{
		Status:  "success",
		Message: "Service registered successfully",
		Data:    service,
	})
}

// GetServices returns all registered services
func (sr *ServiceRegistry) GetServices(w http.ResponseWriter, r *http.Request) {
	services := make([]Service, 0, len(sr.Services))
	for _, service := range sr.Services {
		services = append(services, service)
	}

	respondWithJSON(w, http.StatusOK, Response{
		Status: "success",
		Data:   services,
	})
}

// HealthCheck handles service health check
func (sr *ServiceRegistry) HealthCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceName := vars["name"]

	service, exists := sr.Services[serviceName]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, Response{
			Status:  "error",
			Message: "Service not found",
		})
		return
	}

	// Update last ping
	service.LastPing = time.Now()
	sr.Services[serviceName] = service

	respondWithJSON(w, http.StatusOK, Response{
		Status:  "success",
		Message: "Service is healthy",
		Data:    service,
	})
}

// ServiceClient represents a client for inter-service communication
type ServiceClient struct {
	registry *ServiceRegistry
	client   *http.Client
}

// NewServiceClient creates a new service client
func NewServiceClient(registry *ServiceRegistry) *ServiceClient {
	return &ServiceClient{
		registry: registry,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// CallService makes a request to another service
func (sc *ServiceClient) CallService(serviceName string, path string) (*http.Response, error) {
	service, exists := sc.registry.Services[serviceName]
	if !exists {
		return nil, fmt.Errorf("service %s not found", serviceName)
	}

	url := service.URL + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return sc.client.Do(req)
}

