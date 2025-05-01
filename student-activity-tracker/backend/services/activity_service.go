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
	return &ActivityService{
		collection: collection,
	}
}

// CreateActivity creates a new activity
func (s *ActivityService) CreateActivity(ctx context.Context, activity *models.Activity) error {
	activity.ID = primitive.NewObjectID()
	activity.Status = "planned"
	activity.CreatedAt = time.Now()
	activity.UpdatedAt = time.Now()

	_, err := s.collection.InsertOne(ctx, activity)
	return err
}

// GetActivities retrieves all activities
func (s *ActivityService) GetActivities(ctx context.Context) ([]models.Activity, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var activities []models.Activity
	if err := cursor.All(ctx, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}

// GetActivity retrieves a single activity by ID
func (s *ActivityService) GetActivity(ctx context.Context, id primitive.ObjectID) (*models.Activity, error) {
	var activity models.Activity
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&activity)
	if err != nil {
		return nil, err
