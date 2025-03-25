package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// upgrader upgrades HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{}
// clients holds all currently connected WebSocket clients.
var clients = make(map[*websocket.Conn]bool)
// broadcast is a channel for broadcasting messages to all clients.
var broadcast = make(chan Message)
// mu protects the clients map for concurrent access.
var mu sync.Mutex

// Message represents a chat message sent by a user.
type Message struct {
	Username string `json:"username"` // Name of the sender
	Content  string `json:"content"`  // Message content
}

// handleConnections upgrades HTTP requests to WebSocket connections and reads messages from clients.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// Register new client
	mu.Lock()
	clients[ws] = true
	mu.Unlock()
	for {
