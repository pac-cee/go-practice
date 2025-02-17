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
