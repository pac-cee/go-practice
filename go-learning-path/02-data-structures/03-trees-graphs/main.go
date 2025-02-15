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
