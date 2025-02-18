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
    for i, c := range n.Children {
        if c == child {
            n.Children = append(n.Children[:i], n.Children[i+1:]...)
            return
        }
    }
}
```

## Binary Trees

### Binary Tree Node
```go
type BinaryNode struct {
    Value       interface{}
    Left, Right *BinaryNode
}

// Generic binary node
type BinaryNode[T any] struct {
    Value       T
    Left, Right *BinaryNode[T]
}
```

### Traversal Methods
```go
// Pre-order traversal
func PreOrder(node *BinaryNode) {
    if node == nil {
        return
    }
    fmt.Println(node.Value)
    PreOrder(node.Left)
    PreOrder(node.Right)
}

// In-order traversal
func InOrder(node *BinaryNode) {
    if node == nil {
        return
    }
    InOrder(node.Left)
    fmt.Println(node.Value)
    InOrder(node.Right)
}

// Post-order traversal
func PostOrder(node *BinaryNode) {
    if node == nil {
        return
    }
    PostOrder(node.Left)
    PostOrder(node.Right)
    fmt.Println(node.Value)
}

