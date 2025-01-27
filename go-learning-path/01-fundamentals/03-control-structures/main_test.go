package main

import (
	"strings"
	"testing"
	"time"
)

// TestUserRoles tests the UserRole constants
func TestUserRoles(t *testing.T) {
	// Test role ordering
	if Guest >= User || User >= Moderator || Moderator >= Admin {
		t.Error("User roles are not properly ordered")
	}

	// Test role values
	roles := []UserRole{Guest, User, Moderator, Admin}
	for i, role := range roles {
		if int(role) != i {
			t.Errorf("Expected role value %d, got %d", i, role)
		}
	}
}

// TestCheckUserAccess tests the checkUserAccess function
func TestCheckUserAccess(t *testing.T) {
	tests := []struct {
		name           string
		user          User
		resource      string
		expectedParts []string
	}{
		{
			name: "admin access",
			user: User{
				Name:     "Alice",
				Role:     Admin,
				JoinDate: time.Now(),
