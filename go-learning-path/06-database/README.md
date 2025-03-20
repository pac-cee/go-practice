# Database Integration in Go 🗄️

## 📖 Table of Contents
1. [Database Basics](#database-basics)
2. [SQL Databases](#sql-databases)
3. [NoSQL Databases](#nosql-databases)
4. [ORM Usage](#orm-usage)
5. [Migration Management](#migration-management)
6. [Best Practices](#best-practices)
7. [Exercises](#exercises)

## Database Basics

### Connection Management
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

// Set connection pool settings
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)
```

### Connection Pool Best Practices
1. Set appropriate pool size
2. Monitor connection usage
3. Use connection timeouts
4. Handle connection errors gracefully

## SQL Databases

### Basic CRUD Operations
```go
// Create
result, err := db.Exec(
    "INSERT INTO users (name, email) VALUES ($1, $2)",
    "John Doe",
    "john@example.com",
)

// Read
rows, err := db.Query("SELECT * FROM users WHERE active = $1", true)
for rows.Next() {
    var user User
    err := rows.Scan(&user.ID, &user.Name, &user.Email)
}

// Update
result, err := db.Exec(
    "UPDATE users SET name = $1 WHERE id = $2",
    "Jane Doe",
    1,
)

// Delete
result, err := db.Exec("DELETE FROM users WHERE id = $1", 1)
```

### Prepared Statements
```go
stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES($1, $2)")
defer stmt.Close()

for _, user := range users {
    _, err := stmt.Exec(user.Name, user.Email)
}
```

### Transactions
```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback()

// Perform multiple operations
_, err = tx.Exec(query1, args1...)
_, err = tx.Exec(query2, args2...)

return tx.Commit()
```

## NoSQL Databases

### MongoDB Example
```go
import "go.mongodb.org/mongo-driver/mongo"

client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
defer client.Disconnect(ctx)

collection := client.Database("test").Collection("users")

// Insert
result, err := collection.InsertOne(ctx, bson.D{
    {"name", "John"},
    {"email", "john@example.com"},
})

// Find
var user User
err = collection.FindOne(ctx, bson.D{{"name", "John"}}).Decode(&user)
```

### Redis Example
```go
import "github.com/go-redis/redis/v8"

rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// Set value
err := rdb.Set(ctx, "key", "value", 0).Err()

// Get value
val, err := rdb.Get(ctx, "key").Result()
```

## ORM Usage

### GORM Example
```go
import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string
}

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// Create
user := User{Name: "John", Email: "john@example.com"}
db.Create(&user)

// Read
var user User
db.First(&user, 1)
db.Where("name = ?", "John").First(&user)

// Update
db.Model(&user).Update("Name", "Jane")

// Delete
db.Delete(&user)
```

## Migration Management

### Using golang-migrate
```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create migration
migrate create -ext sql -dir migrations -seq create_users_table

# Run migrations
migrate -path migrations -database "postgres://localhost:5432/dbname?sslmode=disable" up
