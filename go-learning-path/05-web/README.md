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

### HTTPS Support
```go
// Generate certificates
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
```

## Routing

### Basic Router
```go
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api/", apiHandler)
```

### Using Gorilla Mux
```go
router := mux.NewRouter()

router.HandleFunc("/users/{id:[0-9]+}", userHandler).Methods("GET")
router.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler).
    Methods("GET").
    Schemes("https")
```

### Route Groups
```go
api := router.PathPrefix("/api/v1").Subrouter()
api.HandleFunc("/users", usersHandler)
api.HandleFunc("/products", productsHandler)
```

## Middleware

### Logging Middleware
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

### Authentication Middleware
```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if !validateToken(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### Middleware Chain
```go
router.Use(loggingMiddleware)
router.Use(authMiddleware)
router.Use(corsMiddleware)
```

## Templates

### Basic Template
```go
tmpl := template.Must(template.ParseFiles("layout.html"))
tmpl.Execute(w, data)
```

### Template with Layout
```html
<!-- layout.html -->
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    {{template "content" .}}
</body>
</html>

<!-- content.html -->
{{define "content"}}
<h1>{{.Heading}}</h1>
<p>{{.Content}}</p>
{{end}}
```

### Template Functions
```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
    "formatDate": func(t time.Time) string {
        return t.Format("2006-01-02")
    },
}

tmpl := template.New("").Funcs(funcMap)
```

## Database Integration

### Database Connection
```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### CRUD Operations
```go
// Create
stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES($1, $2)")
result, err := stmt.Exec("John", "john@example.com")

// Read
rows, err := db.Query("SELECT * FROM users WHERE active = $1", true)
for rows.Next() {
    var user User
    err := rows.Scan(&user.ID, &user.Name, &user.Email)
}

// Update
stmt, err := db.Prepare("UPDATE users SET name = $1 WHERE id = $2")
result, err := stmt.Exec("Jane", 1)

// Delete
stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
result, err := stmt.Exec(1)
```

## Authentication

### JWT Authentication
```go
func createToken(user *User) (string, error) {
