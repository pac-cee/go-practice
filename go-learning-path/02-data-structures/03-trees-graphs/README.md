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

// Level-order traversal
func LevelOrder(root *BinaryNode) {
    if root == nil {
        return
    }
    queue := []*BinaryNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Println(node.Value)
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

## Binary Search Trees

### BST Node
```go
type BSTNode struct {
    Value       int
    Left, Right *BSTNode
}

// Insert value
func (n *BSTNode) Insert(value int) {
    if value <= n.Value {
        if n.Left == nil {
            n.Left = &BSTNode{Value: value}
        } else {
            n.Left.Insert(value)
        }
    } else {
        if n.Right == nil {
            n.Right = &BSTNode{Value: value}
        } else {
            n.Right.Insert(value)
        }
    }
}

// Search value
func (n *BSTNode) Search(value int) bool {
    if n == nil {
        return false
    }
    if n.Value == value {
        return true
    }
    if value < n.Value {
        return n.Left.Search(value)
    }
