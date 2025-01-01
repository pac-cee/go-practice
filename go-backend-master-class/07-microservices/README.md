# Go Microservices Implementation

This module demonstrates how to implement a microservices architecture in Go, including service discovery, health checks, and circuit breaker patterns.

## Code Explanation

The `main.go` file implements a complete microservices system with the following components:

### 1. Service Model
```go
type Service struct {
    Name     string    `json:"name"`
    URL      string    `json:"url"`
    Status   string    `json:"status"`
    LastPing time.Time `json:"last_ping"`
}

type ServiceRegistry struct {
    Services map[string]Service
}
```
- Service definition
- Registry for service discovery
- Health status tracking
- Last ping monitoring

### 2. Service Registry
```go
func (sr *ServiceRegistry) RegisterService(w http.ResponseWriter, r *http.Request)
func (sr *ServiceRegistry) GetServices(w http.ResponseWriter, r *http.Request)
func (sr *ServiceRegistry) HealthCheck(w http.ResponseWriter, r *http.Request)
```
- Service registration
- Service discovery
- Health monitoring
- Status updates

### 3. Service Client
```go
type ServiceClient struct {
    registry *ServiceRegistry
    client   *http.Client
}
```
- Inter-service communication
- HTTP client management
- Service discovery integration
- Error handling

### 4. Circuit Breaker
```go
type CircuitBreaker struct {
    failures  int
    threshold int
    timeout   time.Duration
    lastError time.Time
}
```
- Failure threshold tracking
- Timeout management
- Error tracking
- Circuit state management

### 5. Circuit Breaker Implementation
```go
func (cb *CircuitBreaker) Execute(fn func() error) error
```
- Failure counting
- Circuit state transitions
- Timeout handling
- Error propagation

## Microservices Patterns

### 1. Service Discovery
- Dynamic service registration
- Health-based service filtering
- Last ping tracking
- Service status updates

### 2. Health Checks
- Regular health monitoring
- Status updates
- Ping timestamp tracking
- Failure detection

### 3. Circuit Breaker
- Prevents cascade failures
- Automatic recovery
- Configurable thresholds
- Timeout management

### 4. Inter-Service Communication
- HTTP-based communication
- Error handling
- Retry logic
- Circuit breaker protection

## API Endpoints

### Service Registry
- `POST /services`: Register a new service
  ```json
  {
    "name": "user-service",
    "url": "http://localhost:8081"
  }
  ```

- `GET /services`: List all services
- `GET /services/{name}/health`: Check service health

## Best Practices Demonstrated
1. Service isolation
