package main

import (
	"expense-tracker/backend-vanilla/config"
	"expense-tracker/backend-vanilla/controllers"
	"expense-tracker/backend-vanilla/middleware"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
