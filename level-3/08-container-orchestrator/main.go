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
