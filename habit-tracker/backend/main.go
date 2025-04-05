package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	r := gin.Default()

	// Public routes
	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(AuthMiddleware())
	{
