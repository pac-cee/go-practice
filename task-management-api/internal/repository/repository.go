package repository

import (
	"github.com/task-management-api/internal/models"
)

// Repository interface defines all database operations
type Repository interface {
	// Task operations
	CreateTask(task *models.Task) error
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
	ListTasks(userID uint) ([]models.Task, error)

	// User operations
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
