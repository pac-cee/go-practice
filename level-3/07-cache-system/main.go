package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type CacheEntry struct {
	Value      interface{} `json:"value"`
	ExpiresAt  time.Time   `json:"expires_at"`
	ReplicaIDs []string    `json:"replica_ids,omitempty"`
}

type Message struct {
	Type    string      `json:"type"`
	Key     string      `json:"key,omitempty"`
	Value   interface{} `json:"value,omitempty"`
	TTL     int         `json:"ttl,omitempty"`
	NodeID  string      `json:"node_id,omitempty"`
	Success bool        `json:"success,omitempty"`
}

type CacheNode struct {
	ID            string
	cache         map[string]CacheEntry
	peers         map[string]net.Conn
	listener      net.Listener
	mutex         sync.RWMutex
	replicaCount  int
	cleanupTicker *time.Ticker
}

func NewCacheNode(id string, port string, replicaCount int) (*CacheNode, error) {
