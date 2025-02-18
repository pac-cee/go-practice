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

// TestAVLTree tests AVL tree operations
func TestAVLTree(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		avl := NewAVL()
		values := []int{5, 3, 7, 2, 4, 6, 8}

		// Test Insert and balance
		for _, v := range values {
			avl.Insert(v)
			if !isBalanced(avl.Root) {
				t.Errorf("Tree not balanced after inserting %d", v)
			}
		}

		// Test height property
		if height := avl.Root.height(); height > 4 {
			t.Errorf("Tree height = %d; want <= 4", height)
		}
	})

	t.Run("rotations", func(t *testing.T) {
		avl := NewAVL()
		
		// Test Left-Left case
		values := []int{30, 20, 10}
		for _, v := range values {
			avl.Insert(v)
		}
		if avl.Root.Value != 20 {
			t.Errorf("Root value after LL rotation = %d; want 20", avl.Root.Value)
		}

		// Test Right-Right case
		avl = NewAVL()
		values = []int{10, 20, 30}
		for _, v := range values {
			avl.Insert(v)
