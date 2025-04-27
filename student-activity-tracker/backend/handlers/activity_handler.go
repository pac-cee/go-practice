package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pac-cee/student-activity-tracker/models"
	"github.com/pac-cee/student-activity-tracker/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActivityHandler handles HTTP requests for activities
type ActivityHandler struct {
	service *services.ActivityService
}

// NewActivityHandler creates a new ActivityHandler
func NewActivityHandler(service *services.ActivityService) *ActivityHandler {
	return &ActivityHandler{
