package services

import (
	"context"
	"time"

	"github.com/pac-cee/student-activity-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ActivityService handles business logic for activities
type ActivityService struct {
	collection *mongo.Collection
}

// NewActivityService creates a new ActivityService
func NewActivityService(collection *mongo.Collection) *ActivityService {
