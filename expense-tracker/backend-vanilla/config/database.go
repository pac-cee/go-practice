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
