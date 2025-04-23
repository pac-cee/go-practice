# Container Orchestration Tool

A simple container orchestration system built in Go, demonstrating container management and monitoring capabilities.

## Features

- Container lifecycle management (create, start, stop, remove)
- Container monitoring (CPU, memory usage)
- REST API interface
- Docker integration
- Real-time container statistics
- Environment variable support
- Port mapping
- Container health checks

## Prerequisites

1. Go 1.19 or later
2. Docker Engine installed and running
3. Docker API access

## API Endpoints

### Create Container
```http
POST /containers
Content-Type: application/json

{
    "name": "my-container",
    "image": "nginx:latest",
    "env": {
        "KEY": "value"
    },
    "ports": {
        "80/tcp": "8080"
    }
}
```

### List Containers
```http
GET /containers
```

### Stop Container
```http
POST /containers/{id}
```

### Remove Container
```http
DELETE /containers/{id}
```

## Implementation Details

### Container Management
- Uses Docker Engine API
- Supports container configuration
- Handles port mappings
- Manages environment variables
- Implements restart policies

### Monitoring
- Real-time CPU usage tracking
- Memory usage monitoring
- Container health status
- Automatic stats collection

### Data Structures

1. **Container**
```go
type Container struct {
    ID           string
