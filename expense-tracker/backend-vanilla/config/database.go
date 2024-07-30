package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

// DatabaseType represents the type of database being used
type DatabaseType string

const (
	MongoDB    DatabaseType = "mongodb"
	PostgreSQL DatabaseType = "postgresql"
	MySQL      DatabaseType = "mysql"
)

// Database interface defines common methods that all database implementations must support
type Database interface {
	Close() error
	Ping() error
	GetType() DatabaseType
}

// MongoDatabase wraps MongoDB connection
type MongoDatabase struct {
	Client *mongo.Client
	DB     *mongo.Database
