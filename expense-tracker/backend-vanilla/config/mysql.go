package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLDB holds the MySQL database instance
type MySQLDB struct {
	*sql.DB
}

// NewMySQLDB creates a new MySQL database connection
func NewMySQLDB() (*MySQLDB, error) {
	// Get database connection details from environment variables
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")

	// Create connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname)

	// Open database connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mysql database: %v", err)
	}

	// Test the connection
	err = db.Ping()
