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
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func validateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
}
```

### Session Management
```go
import "github.com/gorilla/sessions"

var store = sessions.NewCookieStore([]byte("secret-key"))

func sessionHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session-name")
    session.Values["user_id"] = 123
    session.Save(r, w)
}
```

## RESTful APIs

### API Structure
```go
type Server struct {
    router *mux.Router
    db     *sql.DB
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    // Handle request
}
```

### JSON Response
```go
func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
    respondJSON(w, code, map[string]string{"error": message})
}
```

### Request Validation
```go
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

func validateRequest(v interface{}) error {
    validate := validator.New()
    return validate.Struct(v)
}
```

## WebSocket

### WebSocket Server
```go
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // Adjust for production
    },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()
