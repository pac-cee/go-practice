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

type BudgetHandler struct {
	collection *mongo.Collection
}

func NewBudgetHandler(db *mongo.Database) *BudgetHandler {
	return &BudgetHandler{
		collection: db.Collection("budgets"),
	}
}

func (h *BudgetHandler) GetBudgets(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var budgets []models.Budget
