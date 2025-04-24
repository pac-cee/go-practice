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
    Name         string
    Image        string
    Status       string
    CreatedAt    time.Time
    Health       string
    CPU          float64
    Memory       int64
    RestartCount int
}
```

2. **ContainerRequest**
```go
type ContainerRequest struct {
    Name  string
    Image string
    Env   map[string]string
    Ports map[string]string
}
```

## How to Run

1. Start the orchestrator:
   ```bash
   go run main.go
   ```

2. Create a container:
   ```bash
   curl -X POST http://localhost:8080/containers \
     -H "Content-Type: application/json" \
     -d '{
       "name": "web-server",
       "image": "nginx:latest",
       "ports": {
         "80/tcp": "8080"
       }
