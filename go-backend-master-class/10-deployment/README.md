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
