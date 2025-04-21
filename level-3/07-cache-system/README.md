# Distributed Cache System

A distributed in-memory cache system built in Go, featuring peer-to-peer replication and automatic data expiration.

## Features

- In-memory key-value storage
- Distributed architecture with peer-to-peer replication
- TTL (Time-To-Live) support for cache entries
- Automatic cleanup of expired entries
- TCP-based communication between nodes
- Thread-safe operations
- JSON-based protocol for inter-node communication

## Architecture

The system implements a distributed cache with the following components:

1. **Cache Node**
   - Stores key-value pairs in memory
   - Handles peer connections
   - Manages data replication
   - Performs automatic cleanup

2. **Cache Entry**
   - Stores value with expiration time
   - Tracks replica locations

3. **Message Protocol**
   - Set operations
   - Get operations
   - Delete operations
   - Node join/leave operations

## Prerequisites

- Go 1.19 or later

## How to Run

1. Start the first node:
   ```bash
   go run main.go node1 8081
   ```

2. Start additional nodes:
   ```bash
   go run main.go node2 8082 localhost:8081
   go run main.go node3 8083 localhost:8081 localhost:8082
   ```

## Protocol Specification

### Message Types

1. **Set**
   ```json
   {
     "type": "set",
     "key": "example-key",
     "value": "example-value",
     "ttl": 3600,
     "node_id": "node1"
   }
   ```

2. **Get**
   ```json
   {
     "type": "get",
     "key": "example-key"
   }
   ```

3. **Delete**
   ```json
