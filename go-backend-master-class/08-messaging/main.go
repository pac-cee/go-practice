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
