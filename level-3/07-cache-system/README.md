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
