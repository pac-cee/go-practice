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

