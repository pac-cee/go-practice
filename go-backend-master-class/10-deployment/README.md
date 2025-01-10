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

