package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	username string
}

type Message struct {
	Type     string `json:"type"`
	Content  string `json:"content"`
	Username string `json:"username"`
	Time     string `json:"time"`
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
