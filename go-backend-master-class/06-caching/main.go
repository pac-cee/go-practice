package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// Product represents a product in our system
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
