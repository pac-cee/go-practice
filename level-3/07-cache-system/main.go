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
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	node := &CacheNode{
		ID:           id,
		cache:        make(map[string]CacheEntry),
		peers:        make(map[string]net.Conn),
		listener:     listener,
		replicaCount: replicaCount,
	}

	// Start cleanup routine
	node.cleanupTicker = time.NewTicker(1 * time.Minute)
	go node.cleanup()

	return node, nil
}

func (n *CacheNode) Set(key string, value interface{}, ttl int) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	expiresAt := time.Now().Add(time.Duration(ttl) * time.Second)
	entry := CacheEntry{
		Value:     value,
		ExpiresAt: expiresAt,
	}

	n.cache[key] = entry

	// Replicate to peers
	msg := Message{
		Type:   "set",
		Key:    key,
		Value:  value,
		TTL:    ttl,
		NodeID: n.ID,
	}

	return n.replicateToPeers(msg)
}

func (n *CacheNode) Get(key string) (interface{}, bool) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()

	entry, exists := n.cache[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(entry.ExpiresAt) {
		delete(n.cache, key)
		return nil, false
	}
