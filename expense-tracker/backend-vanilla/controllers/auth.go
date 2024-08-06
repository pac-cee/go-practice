package controllers

import (
	"context"
	"expense-tracker/backend-vanilla/config"
	"expense-tracker/backend-vanilla/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Password validation rules
var passwordRules = struct {
	MinLength      int
	MaxLength      int
	RequireUpper   bool
	RequireLower   bool
	RequireNumber  bool
	RequireSpecial bool
	AllowedSpecial string
}{
	MinLength:      8,
	MaxLength:      32,
	RequireUpper:   true,
	RequireLower:   true,
	RequireNumber:  true,
	RequireSpecial: true,
	AllowedSpecial: "!@#$%^&*(),.?\":{}|<>",
}

func validatePassword(password string) error {
	if len(password) < passwordRules.MinLength {
		return fmt.Errorf("password must be at least %d characters long", passwordRules.MinLength)
	}

	if len(password) > passwordRules.MaxLength {
		return fmt.Errorf("password must be less than %d characters", passwordRules.MaxLength)
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case strings.ContainsRune(passwordRules.AllowedSpecial, char):
			hasSpecial = true
		}
	}

	if passwordRules.RequireUpper && !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if passwordRules.RequireLower && !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	if passwordRules.RequireNumber && !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}

	if passwordRules.RequireSpecial && !hasSpecial {
		return fmt.Errorf("password must contain at least one special character (%s)", passwordRules.AllowedSpecial)
	}

	return nil
}

func Register(c *gin.Context) {
	var input struct {
		Name     string  `json:"name" binding:"required"`
		Email    string  `json:"email" binding:"required,email"`
		Password string  `json:"password" binding:"required"`
		Limit    float64 `json:"spending_limit"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate password
