# Go Authentication

This module demonstrates how to implement secure authentication in Go using JWT (JSON Web Tokens) and bcrypt for password hashing.

## Code Explanation

The `main.go` file implements a complete authentication system with the following components:

### 1. Data Models
```go
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```
- User model for authentication
- JWT claims structure
- Standardized API response

### 2. User Registration
```go
func RegisterHandler(w http.ResponseWriter, r *http.Request)
```
- Validates user input
- Checks for existing users
- Securely hashes passwords using bcrypt
- Stores user credentials
- Handles registration errors

### 3. User Login
```go
func LoginHandler(w http.ResponseWriter, r *http.Request)
```
- Validates credentials
- Compares password hashes
- Generates JWT token
- Sets token expiration
- Returns authentication token

### 4. Authentication Middleware
```go
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc
```
- Validates JWT tokens
- Extracts user claims
- Handles token expiration
- Protects routes
- Maintains security context

### 5. Protected Routes
```go
func ProtectedHandler(w http.ResponseWriter, r *http.Request)
```
- Example of protected endpoint
- Requires valid JWT
- Demonstrates middleware usage
- Shows secure access pattern

## Security Features

### Password Security
1. Bcrypt hashing
2. Secure password comparison
3. Protection against timing attacks
4. Configurable hash cost
