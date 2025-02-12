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
		cache := NewCache()

		// Test Set and Get
		cache.Set("key1", "value1", time.Minute)
		if value, exists := cache.Get("key1"); !exists || value != "value1" {
			t.Errorf("cache.Get(key1) = %v, %v; want value1, true", value, exists)
		}

		// Test non-existent key
		if _, exists := cache.Get("key2"); exists {
			t.Error("cache should not contain key2")
		}

		// Test Delete
		cache.Delete("key1")
		if _, exists := cache.Get("key1"); exists {
			t.Error("cache should not contain key1 after deletion")
		}
	})

	t.Run("expiration", func(t *testing.T) {
		cache := NewCache()

		// Add item with short TTL
		cache.Set("key1", "value1", time.Millisecond*100)

		// Wait for expiration
		time.Sleep(time.Millisecond * 200)

		// Check if item has expired
		if _, exists := cache.Get("key1"); exists {
			t.Error("cache item should have expired")
		}
	})
}

// TestFrequencyCounter tests the FrequencyCounter
func TestFrequencyCounter(t *testing.T) {
	t.Run("basic counting", func(t *testing.T) {
		counter := NewFrequencyCounter()

		// Add words
		words := []string{"apple", "banana", "apple", "cherry", "apple", "date"}
		for _, word := range words {
			counter.Add(word)
		}

		// Test counts
		tests := []struct {
			word     string
			expected int
		}{
			{"apple", 3},
			{"banana", 1},
			{"cherry", 1},
			{"date", 1},
			{"fig", 0}, // non-existent word
		}

		for _, tt := range tests {
			if count := counter.Count(tt.word); count != tt.expected {
				t.Errorf("Count(%s) = %d; want %d", tt.word, count, tt.expected)
			}
		}
	})

	t.Run("top N", func(t *testing.T) {
		counter := NewFrequencyCounter()

		// Add words with different frequencies
		words := []string{
			"apple", "apple", "apple",
			"banana", "banana",
			"cherry",
		}
		for _, word := range words {
