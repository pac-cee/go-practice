# Go Database Integration

This module demonstrates how to integrate a PostgreSQL database with Go using the standard `database/sql` package and the `lib/pq` driver.

## Code Explanation

The `main.go` file implements a complete database layer with the following components:

### 1. Data Model
```go
type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
}
```
- Defines the User model structure
- Maps to the database table columns
- Uses appropriate Go types for each field

### 2. Database Configuration
```go
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "gobackend"
)
```
- Centralizes database connection parameters
- Makes configuration easily modifiable
- Follows standard PostgreSQL connection format

### 3. Repository Pattern
```go
type UserRepository struct {
