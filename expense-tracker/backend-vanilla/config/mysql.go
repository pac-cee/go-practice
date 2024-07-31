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
