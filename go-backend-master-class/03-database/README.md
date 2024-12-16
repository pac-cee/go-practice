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
    db *sql.DB
}
```
- Implements the repository pattern for data access
- Encapsulates all database operations
- Provides clean separation of concerns

### 4. Database Initialization
```go
func initDB() (*sql.DB, error)
```
- Establishes database connection
- Uses connection pooling by default
- Verifies connection with `db.Ping()`
- Returns configured database instance

### 5. Schema Management
```go
func (r *UserRepository) createTables() error
```
- Creates necessary database tables
- Uses `IF NOT EXISTS` for idempotency
- Defines proper constraints (UNIQUE, NOT NULL)
- Sets up automatic timestamps

### 6. CRUD Operations
The repository implements full CRUD (Create, Read, Update, Delete) operations:

#### Create
```go
func (r *UserRepository) CreateUser(username, email string) (*User, error)
```
- Uses parameterized queries for safety
- Returns complete user object with ID
- Handles database errors properly

#### Read
```go
