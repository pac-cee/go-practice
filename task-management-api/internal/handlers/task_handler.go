package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/task-management-api/internal/models"
	"github.com/task-management-api/internal/repository"
)

type TaskHandler struct {
	repo repository.Repository
}

func NewTaskHandler(repo repository.Repository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

// CreateTask handles task creation
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Get userID from JWT token
	task.UserID = 1 // Temporary

	if err := h.repo.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
