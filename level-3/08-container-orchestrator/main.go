package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

type Container struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Health      string    `json:"health"`
	CPU         float64   `json:"cpu_usage"`
	Memory      int64     `json:"memory_usage"`
	RestartCount int      `json:"restart_count"`
}

type ContainerRequest struct {
	Name  string            `json:"name"`
	Image string            `json:"image"`
	Env   map[string]string `json:"env"`
	Ports map[string]string `json:"ports"`
}

type Orchestrator struct {
	client     *client.Client
	containers map[string]*Container
	mutex      sync.RWMutex
}

func NewOrchestrator() (*Orchestrator, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &Orchestrator{
		client:     cli,
		containers: make(map[string]*Container),
	}, nil
}

func (o *Orchestrator) CreateContainer(w http.ResponseWriter, r *http.Request) {
	var req ContainerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare environment variables
	var env []string
	for k, v := range req.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	// Prepare port bindings
	portBindings := make(map[string][]types.PortBinding)
	exposedPorts := make(map[string]struct{})
	for containerPort, hostPort := range req.Ports {
		portBindings[containerPort] = []types.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		}
		exposedPorts[containerPort] = struct{}{}
	}

	// Create container
	resp, err := o.client.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        req.Image,
			Env:          env,
			ExposedPorts: exposedPorts,
		},
		&container.HostConfig{
			PortBindings: portBindings,
			RestartPolicy: container.RestartPolicy{
				Name: "unless-stopped",
			},
		},
		nil,
		nil,
		req.Name,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Start container
	if err := o.client.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	container := &Container{
		ID:        resp.ID,
		Name:      req.Name,
		Image:     req.Image,
