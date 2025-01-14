# Go Deployment and Monitoring

This module demonstrates how to deploy a Go application using Docker and Kubernetes, with built-in monitoring using Prometheus metrics.

## Code Explanation

### 1. Application Components

#### Metrics Implementation
```go
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint"},
    )
)
```
- Request counting
- Duration tracking
- Label-based metrics
- Prometheus integration

#### Middleware
```go
func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc
```
- Request metrics
- Response timing
- Logging
- Error tracking

#### Health Checks
```go
func healthHandler(w http.ResponseWriter, r *http.Request)
func readinessHandler(w http.ResponseWriter, r *http.Request)
```
- Liveness probe
- Readiness probe
- Version info
- Uptime tracking

### 2. Docker Configuration

```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```
- Multi-stage build
- Minimal final image
- Dependency management
- Port configuration

### 3. Kubernetes Deployment

#### Deployment Configuration
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-backend-app
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: go-backend
        image: go-backend:latest
        resources:
          limits:
            cpu: "500m"
            memory: "512Mi"
```
- Replica management
- Resource limits
- Container configuration
- Health probes

#### Service Configuration
```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-backend-service
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 8080
```
- Load balancing
- Port mapping
- Service discovery
- External access

## Deployment Process

### 1. Building the Docker Image
```bash
# Build the image
docker build -t go-backend:latest .

# Run locally
docker run -p 8080:8080 go-backend:latest
```

### 2. Kubernetes Deployment
```bash
# Apply the deployment
kubectl apply -f k8s/deployment.yaml

