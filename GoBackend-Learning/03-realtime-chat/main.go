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
