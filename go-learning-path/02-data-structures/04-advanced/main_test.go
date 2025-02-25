package main

import (
	"reflect"
	"testing"
)

// TestMinHeap tests heap operations
func TestMinHeap(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		heap := NewMinHeap()
		values := []int{5, 3, 7, 1, 4}

		// Test Insert
		for _, v := range values {
			heap.Insert(v)
		}

		// Test ExtractMin
		expected := []int{1, 3, 4, 5, 7}
		for _, want := range expected {
			got, err := heap.ExtractMin()
			if err != nil {
				t.Errorf("ExtractMin() error = %v", err)
			}
			if got != want {
				t.Errorf("ExtractMin() = %v, want %v", got, want)
			}
		}

		// Test empty heap
		_, err := heap.ExtractMin()
		if err == nil {
			t.Error("ExtractMin() on empty heap should return error")
		}
	})
}

// TestTrie tests trie operations
func TestTrie(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		trie := NewTrie()
		words := []string{"apple", "app", "apricot", "banana"}

		// Test Insert and Search
		for _, word := range words {
			trie.Insert(word)
		}

		tests := []struct {
			word     string
			exists   bool
			message  string
		}{
			{"apple", true, "existing word"},
			{"app", true, "existing word"},
			{"apt", false, "non-existing word"},
			{"", true, "empty string"},
		}

		for _, tt := range tests {
			t.Run(tt.message, func(t *testing.T) {
				if got := trie.Search(tt.word); got != tt.exists {
					t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.exists)
				}
			})
		}

		// Test StartsWith
		prefixes := []struct {
			prefix   string
			exists   bool
			message  string
		}{
			{"ap", true, "existing prefix"},
			{"ban", true, "existing prefix"},
