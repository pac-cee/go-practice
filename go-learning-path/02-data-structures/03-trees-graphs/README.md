# Trees and Graphs in Go 🌳

## 📖 Table of Contents
1. [Trees](#trees)
2. [Binary Trees](#binary-trees)
3. [Binary Search Trees](#binary-search-trees)
4. [Graphs](#graphs)
5. [Common Algorithms](#common-algorithms)
6. [Best Practices](#best-practices)
7. [Exercises](#exercises)

## Trees

### Tree Node
```go
type Node struct {
    Value    interface{}
    Children []*Node
}

// Generic tree node
type TreeNode[T any] struct {
    Value    T
    Children []*TreeNode[T]
}
```

### Tree Operations
```go
// Add child
func (n *Node) AddChild(value interface{}) *Node {
    child := &Node{Value: value}
    n.Children = append(n.Children, child)
    return child
}

// Remove child
func (n *Node) RemoveChild(child *Node) {
