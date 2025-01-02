package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Message represents a message in our system
type Message struct {
	ID        string      `json:"id"`
	Topic     string      `json:"topic"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

// Subscription represents a subscription to a topic
type Subscription struct {
	Topic    string
	Channel  chan Message
	ClientID string
}

// MessageBroker handles pub/sub messaging
type MessageBroker struct {
	sync.RWMutex
	subscriptions map[string][]Subscription
	topics        map[string][]Message
}

// WebSocketClient represents a connected websocket client
type WebSocketClient struct {
	ID       string
	Conn     *websocket.Conn
	Broker   *MessageBroker
	Topics   []string
	SendChan chan Message
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for demo
	},
}

// NewMessageBroker creates a new message broker
func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		subscriptions: make(map[string][]Subscription),
