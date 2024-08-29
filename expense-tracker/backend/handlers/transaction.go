package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"expense-tracker/models"
)

type TransactionHandler struct {
	collection *mongo.Collection
}

func NewTransactionHandler(db *mongo.Database) *TransactionHandler {
	return &TransactionHandler{
		collection: db.Collection("transactions"),
	}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse query parameters
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	category := c.Query("category")
	account := c.Query("account")

	// Build filter
	filter := bson.M{}
	if startDate != "" {
		start, err := time.Parse(time.RFC3339, startDate)
		if err == nil {
			filter["date"] = bson.M{"$gte": start}
		}
	}
	if endDate != "" {
		end, err := time.Parse(time.RFC3339, endDate)
		if err == nil {
			if _, ok := filter["date"]; ok {
				filter["date"].(bson.M)["$lte"] = end
			} else {
				filter["date"] = bson.M{"$lte": end}
			}
		}
	}
	if category != "" {
		filter["category"] = category
	}
