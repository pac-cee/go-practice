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
