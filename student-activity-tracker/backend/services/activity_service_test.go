package services

import (
	"context"
	"testing"
	"time"

	"github.com/pac-cee/student-activity-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) (*mongo.Collection, func()) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection := client.Database("test_db").Collection("test_activities")

	return collection, func() {
		collection.Drop(ctx)
		client.Disconnect(ctx)
	}
}

func TestActivityService_CreateActivity(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	service := NewActivityService(collection)
	activity := &models.Activity{
		Title:       "Test Activity",
		Description: "Test Description",
		Duration:    30,
	}

	err := service.CreateActivity(context.Background(), activity)
	if err != nil {
		t.Errorf("Failed to create activity: %v", err)
	}

	if activity.ID.IsZero() {
		t.Error("Activity ID should not be zero")
	}
	if activity.Status != "planned" {
		t.Errorf("Expected status 'planned', got %s", activity.Status)
	}
}

func TestActivityService_GetActivities(t *testing.T) {
	collection, cleanup := setupTestDB(t)
	defer cleanup()

	service := NewActivityService(collection)
	
	// Create test activities
	activities := []models.Activity{
		{
			Title:       "Activity 1",
			Description: "Description 1",
			Duration:    30,
		},
		{
			Title:       "Activity 2",
			Description: "Description 2",
			Duration:    45,
		},
	}

	for _, a := range activities {
		err := service.CreateActivity(context.Background(), &a)
		if err != nil {
