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
