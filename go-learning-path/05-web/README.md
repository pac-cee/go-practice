# Web Development in Go 🌐

## 📖 Table of Contents
1. [HTTP Server](#http-server)
2. [Routing](#routing)
3. [Middleware](#middleware)
4. [Templates](#templates)
5. [Database Integration](#database-integration)
6. [Authentication](#authentication)
7. [RESTful APIs](#restful-apis)
8. [WebSocket](#websocket)
9. [Best Practices](#best-practices)
10. [Exercises](#exercises)

## HTTP Server

### Basic HTTP Server
```go
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
```

### Server Configuration
```go
server := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
    Handler:      router,
}
```
