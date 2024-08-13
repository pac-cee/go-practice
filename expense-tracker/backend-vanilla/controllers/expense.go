package controllers

import (
	"context"
	"expense-tracker/backend-vanilla/config"
	"expense-tracker/backend-vanilla/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateExpense(c *gin.Context) {
	userID := c.MustGet("user_id").(primitive.ObjectID)

	var input struct {
		Amount      float64   `json:"amount" binding:"required"`
		Category    string    `json:"category" binding:"required"`
		Description string    `json:"description"`
		Date        time.Time `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := models.Expense{
		UserID:      userID,
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		Date:        input.Date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
