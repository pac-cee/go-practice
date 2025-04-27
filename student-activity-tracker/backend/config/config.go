package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	MongoURI      string
	DatabaseName  string
	ServerAddress string
	AllowedOrigins []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DatabaseName:  getEnv("DB_NAME", "student_tracker"),
