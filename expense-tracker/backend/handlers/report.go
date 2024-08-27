package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"expense-tracker/models"
)

type ReportHandler struct {
	db *mongo.Database
}

func NewReportHandler(db *mongo.Database) *ReportHandler {
	return &ReportHandler{
		db: db,
	}
}

func (h *ReportHandler) GetTransactionReport(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse date range
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		start = time.Now().AddDate(0, -1, 0) // Default to last month
	}

	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		end = time.Now() // Default to now
	}

	// Build date filter
	dateFilter := bson.M{
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	// Get transactions
	transactionsColl := h.db.Collection("transactions")
	cursor, err := transactionsColl.Find(ctx, dateFilter, options.Find().SetSort(bson.D{{Key: "date", Value: -1}}))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var transactions []models.Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate totals
	var totalIncome, totalExpenses float64
	categoryMap := make(map[string]*models.CategoryExpenseSummary)

	for _, t := range transactions {
		if t.Amount >= 0 {
			totalIncome += t.Amount
		} else {
			totalExpenses += -t.Amount
			// Track category expenses
			if summary, exists := categoryMap[t.Category]; exists {
				summary.Amount += -t.Amount
				summary.Count++
			} else {
				categoryMap[t.Category] = &models.CategoryExpenseSummary{
					Category: t.Category,
					Amount:   -t.Amount,
					Count:    1,
				}
			}
		}
	}

	// Convert category map to slice
	categories := make([]models.CategoryExpenseSummary, 0, len(categoryMap))
	for _, summary := range categoryMap {
		categories = append(categories, *summary)
	}

