package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pac-cee/student-activity-tracker/models"
	"github.com/pac-cee/student-activity-tracker/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestHandler(t *testing.T) (*ActivityHandler, *gin.Engine, func()) {
	// Connect to test database
