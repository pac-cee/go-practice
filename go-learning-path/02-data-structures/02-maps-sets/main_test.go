package main

import (
	"reflect"
	"testing"
	"time"
)

// TestSet tests the Set data structure
func TestSet(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		set := NewSet()

		// Test Add and Contains
		set.Add("apple")
		if !set.Contains("apple") {
			t.Error("set should contain 'apple'")
		}

