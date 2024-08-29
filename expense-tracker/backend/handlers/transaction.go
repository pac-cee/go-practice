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
