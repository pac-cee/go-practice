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
		service: service,
	}
}

// RegisterRoutes registers the routes for activities
func (h *ActivityHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/activities", h.CreateActivity)
		api.GET("/activities", h.GetActivities)
		api.PUT("/activities/:id", h.UpdateActivity)
		api.DELETE("/activities/:id", h.DeleteActivity)
		api.PUT("/activities/:id/start", h.StartActivity)
		api.PUT("/activities/:id/complete", h.CompleteActivity)
	}
}

// CreateActivity handles the creation of a new activity
func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateActivity(c.Request.Context(), &activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

// GetActivities handles retrieving all activities
func (h *ActivityHandler) GetActivities(c *gin.Context) {
	activities, err := h.service.GetActivities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
