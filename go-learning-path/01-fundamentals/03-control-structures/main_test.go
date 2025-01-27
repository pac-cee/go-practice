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
