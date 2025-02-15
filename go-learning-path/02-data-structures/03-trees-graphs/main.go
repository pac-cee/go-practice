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

