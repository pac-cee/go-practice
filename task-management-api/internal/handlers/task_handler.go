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
