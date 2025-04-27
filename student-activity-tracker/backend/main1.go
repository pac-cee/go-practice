package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Activity represents a student's planned activity
type Activity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Duration    int               `bson:"duration" json:"duration"` // in minutes
	Status      string             `bson:"status" json:"status"`    // planned, in-progress, completed
	StartTime   *time.Time         `bson:"startTime" json:"startTime,omitempty"`
	EndTime     *time.Time         `bson:"endTime" json:"endTime,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func init() {
	// Set up MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("student_tracker").Collection("activities")
}

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4321"} // Astro's default port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// Routes
	r.POST("/api/activities", createActivity)
	r.GET("/api/activities", getActivities)
	r.PUT("/api/activities/:id", updateActivity)
	r.DELETE("/api/activities/:id", deleteActivity)
	r.PUT("/api/activities/:id/start", startActivity)
	r.PUT("/api/activities/:id/complete", completeActivity)

	log.Fatal(r.Run(":8080"))
}

func createActivity(c *gin.Context) {
