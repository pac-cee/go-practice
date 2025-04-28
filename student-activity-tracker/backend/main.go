package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pac-cee/student-activity-tracker/config"
	"github.com/pac-cee/student-activity-tracker/database"
	"github.com/pac-cee/student-activity-tracker/handlers"
	"github.com/pac-cee/student-activity-tracker/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.NewMongoDB(cfg.MongoURI, cfg.DatabaseName)
