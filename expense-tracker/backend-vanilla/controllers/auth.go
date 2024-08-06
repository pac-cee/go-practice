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
