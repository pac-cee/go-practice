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
		bst.inOrder(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrder(node.Right, result)
	}
}

// AVLNode represents a node in an AVL tree
type AVLNode struct {
	Value       int
	Left, Right *AVLNode
	Height      int
}

// AVLTree represents an AVL tree
type AVLTree struct {
	Root *AVLNode
}

// NewAVL creates a new AVL tree
func NewAVL() *AVLTree {
	return &AVLTree{}
}

// Height returns the height of a node
func (n *AVLNode) height() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// UpdateHeight updates the height of a node
func (n *AVLNode) updateHeight() {
	n.Height = 1 + max(getHeight(n.Left), getHeight(n.Right))
}

func getHeight(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetBalance returns the balance factor of a node
func (n *AVLNode) getBalance() int {
	if n == nil {
		return 0
	}
	return getHeight(n.Left) - getHeight(n.Right)
}

// RotateRight performs a right rotation
func (avl *AVLTree) rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.updateHeight()
	x.updateHeight()

	return x
}

// RotateLeft performs a left rotation
func (avl *AVLTree) rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.updateHeight()
	y.updateHeight()

	return y
}

// Insert adds a new value to the AVL tree
func (avl *AVLTree) Insert(value int) {
	avl.Root = avl.insertNode(avl.Root, value)
}

func (avl *AVLTree) insertNode(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return &AVLNode{Value: value, Height: 1}
	}
