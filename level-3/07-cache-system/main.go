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
