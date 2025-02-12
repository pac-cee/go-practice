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

	t.Run("set operations", func(t *testing.T) {
		set1 := NewSet()
		set1.Add("apple")
		set1.Add("banana")
		set1.Add("cherry")

		set2 := NewSet()
		set2.Add("banana")
		set2.Add("cherry")
		set2.Add("date")

		// Test Union
		union := set1.Union(set2)
		expectedUnion := []string{"apple", "banana", "cherry", "date"}
		if !sameElements(union.Items(), expectedUnion) {
			t.Errorf("Union = %v; want %v", union.Items(), expectedUnion)
		}

		// Test Intersection
		intersection := set1.Intersection(set2)
		expectedIntersection := []string{"banana", "cherry"}
		if !sameElements(intersection.Items(), expectedIntersection) {
			t.Errorf("Intersection = %v; want %v", intersection.Items(), expectedIntersection)
		}

		// Test Difference
		difference := set1.Difference(set2)
		expectedDifference := []string{"apple"}
		if !sameElements(difference.Items(), expectedDifference) {
			t.Errorf("Difference = %v; want %v", difference.Items(), expectedDifference)
		}
	})
}

// TestCache tests the Cache data structure
func TestCache(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
