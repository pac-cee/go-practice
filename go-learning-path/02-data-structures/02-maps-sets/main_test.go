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

		// Test Size
		if set.Size() != 1 {
			t.Errorf("set size = %d; want 1", set.Size())
		}

		// Test Remove
		set.Remove("apple")
		if set.Contains("apple") {
			t.Error("set should not contain 'apple' after removal")
		}

		// Test Items
		set.Add("banana")
		set.Add("cherry")
		items := set.Items()
		if len(items) != 2 {
			t.Errorf("items length = %d; want 2", len(items))
		}
	})
