// Package main demonstrates trees and graphs in Go
package main

import (
	"fmt"
	"math"
)

// BinaryNode represents a node in a binary tree
type BinaryNode[T any] struct {
	Value       T
	Left, Right *BinaryNode[T]
}

// BinarySearchTree represents a binary search tree
type BinarySearchTree struct {
	Root *BinaryNode[int]
}

// NewBST creates a new binary search tree
func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert adds a new value to the BST
func (bst *BinarySearchTree) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &BinaryNode[int]{Value: value}
		return
	}
	bst.insertNode(bst.Root, value)
}

func (bst *BinarySearchTree) insertNode(node *BinaryNode[int], value int) {
	if value <= node.Value {
		if node.Left == nil {
			node.Left = &BinaryNode[int]{Value: value}
		} else {
			bst.insertNode(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinaryNode[int]{Value: value}
		} else {
			bst.insertNode(node.Right, value)
		}
	}
}

// Search looks for a value in the BST
func (bst *BinarySearchTree) Search(value int) bool {
	return bst.searchNode(bst.Root, value)
}

func (bst *BinarySearchTree) searchNode(node *BinaryNode[int], value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if value < node.Value {
		return bst.searchNode(node.Left, value)
	}
	return bst.searchNode(node.Right, value)
}

// InOrderTraversal performs in-order traversal of the BST
func (bst *BinarySearchTree) InOrderTraversal() []int {
	var result []int
	bst.inOrder(bst.Root, &result)
	return result
}

func (bst *BinarySearchTree) inOrder(node *BinaryNode[int], result *[]int) {
	if node != nil {
