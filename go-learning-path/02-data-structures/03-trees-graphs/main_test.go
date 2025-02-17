package main

import (
	"reflect"
	"testing"
)

// TestBinarySearchTree tests BST operations
func TestBinarySearchTree(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		bst := NewBST()
		values := []int{5, 3, 7, 2, 4, 6, 8}

		// Test Insert
		for _, v := range values {
			bst.Insert(v)
		}

		// Test Search
		for _, v := range values {
			if !bst.Search(v) {
				t.Errorf("Search(%d) = false; want true", v)
			}
		}

		// Test non-existent value
		if bst.Search(9) {
			t.Error("Search(9) = true; want false")
		}

		// Test InOrderTraversal
		expected := []int{2, 3, 4, 5, 6, 7, 8}
		result := bst.InOrderTraversal()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InOrderTraversal() = %v; want %v", result, expected)
		}
	})
}
