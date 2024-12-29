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
