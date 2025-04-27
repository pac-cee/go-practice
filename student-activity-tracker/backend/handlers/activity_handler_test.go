package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pac-cee/student-activity-tracker/models"
	"github.com/pac-cee/student-activity-tracker/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestHandler(t *testing.T) (*ActivityHandler, *gin.Engine, func()) {
	// Connect to test database
	client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection := client.Database("test_db").Collection("test_activities")
	service := services.NewActivityService(collection)
	handler := NewActivityHandler(service)

	// Set up Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler.RegisterRoutes(router)

	return handler, router, func() {
		collection.Drop(nil)
		client.Disconnect(nil)
	}
}
