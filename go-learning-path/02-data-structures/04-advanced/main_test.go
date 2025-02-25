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
			{"cat", false, "non-existing prefix"},
		}

		for _, tt := range prefixes {
			t.Run(tt.message, func(t *testing.T) {
				if got := trie.StartsWith(tt.prefix); got != tt.exists {
					t.Errorf("StartsWith(%q) = %v, want %v", tt.prefix, got, tt.exists)
				}
			})
		}
	})
}

// TestDisjointSet tests union-find operations
func TestDisjointSet(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		ds := NewDisjointSet(5)

		// Test initial state
		for i := 0; i < 5; i++ {
			if got := ds.Find(i); got != i {
				t.Errorf("Find(%d) = %d, want %d", i, got, i)
			}
		}

		// Test Union
		unions := [][2]int{{0, 1}, {2, 3}, {1, 2}}
		for _, u := range unions {
			ds.Union(u[0], u[1])
		}

		// Test connected components
		tests := []struct {
			x, y     int
			connected bool
			message  string
		}{
			{0, 3, true, "should be connected"},
			{1, 4, false, "should not be connected"},
		}

		for _, tt := range tests {
			t.Run(tt.message, func(t *testing.T) {
				if got := ds.Find(tt.x) == ds.Find(tt.y); got != tt.connected {
					t.Errorf("Find(%d) == Find(%d) = %v, want %v", 
						tt.x, tt.y, got, tt.connected)
				}
			})
		}
	})
}

// TestBloomFilter tests bloom filter operations
func TestBloomFilter(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		bf := NewBloomFilter(100, 3)
		items := []string{"apple", "banana", "orange"}

		// Test Add and Contains
		for _, item := range items {
			bf.Add(item)
		}

		// Test existing items
		for _, item := range items {
			if !bf.Contains(item) {
				t.Errorf("Contains(%q) = false, want true", item)
			}
		}

		// Test false positives
		falsePositives := 0
		nonExistingItems := []string{"grape", "melon", "peach"}
		for _, item := range nonExistingItems {
			if bf.Contains(item) {
				falsePositives++
