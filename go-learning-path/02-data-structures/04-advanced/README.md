# Advanced Data Structures in Go 🚀

## 📖 Table of Contents
1. [Priority Queues](#priority-queues)
2. [Trie](#trie)
3. [Union-Find](#union-find)
4. [Skip List](#skip-list)
5. [Bloom Filter](#bloom-filter)
6. [LRU Cache](#lru-cache)
7. [Exercises](#exercises)

## Priority Queues

### Heap Implementation
```go
type Heap struct {
    items    []int
    size     int
    capacity int
}

func NewHeap(capacity int) *Heap {
    return &Heap{
        items:    make([]int, capacity),
        capacity: capacity,
    }
}

func (h *Heap) Parent(i int) int {
    return (i - 1) / 2
}

func (h *Heap) LeftChild(i int) int {
    return 2*i + 1
}

func (h *Heap) RightChild(i int) int {
    return 2*i + 2
}

func (h *Heap) Swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}
```

### Heap Operations
```go
// Insert into min heap
func (h *Heap) Insert(key int) error {
    if h.size == h.capacity {
        return errors.New("heap is full")
    }

    h.size++
    i := h.size - 1
    h.items[i] = key

    // Fix min heap property
    for i != 0 && h.items[h.Parent(i)] > h.items[i] {
        h.Swap(i, h.Parent(i))
        i = h.Parent(i)
    }
    return nil
}

// Extract minimum element
func (h *Heap) ExtractMin() (int, error) {
    if h.size <= 0 {
        return 0, errors.New("heap is empty")
    }
    if h.size == 1 {
        h.size--
        return h.items[0], nil
    }

    root := h.items[0]
    h.items[0] = h.items[h.size-1]
    h.size--
    h.MinHeapify(0)

    return root, nil
}
```

## Trie

### Trie Node
```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{
            children: make(map[rune]*TrieNode),
        },
    }
}
```

### Trie Operations
```go
// Insert a word
func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{
                children: make(map[rune]*TrieNode),
            }
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

// Search for a word
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}

// Search prefix
func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if node.children[ch] == nil {
            return false
        }
        node = node.children[ch]
    }
    return true
}
```

## Union-Find

### Disjoint Set
```go
type DisjointSet struct {
    parent []int
    rank   []int
}

func NewDisjointSet(size int) *DisjointSet {
    parent := make([]int, size)
    rank := make([]int, size)
    for i := range parent {
        parent[i] = i
    }
    return &DisjointSet{
        parent: parent,
        rank:   rank,
    }
}
```

### Union-Find Operations
```go
// Find with path compression
func (ds *DisjointSet) Find(x int) int {
    if ds.parent[x] != x {
        ds.parent[x] = ds.Find(ds.parent[x])
    }
    return ds.parent[x]
}

// Union by rank
func (ds *DisjointSet) Union(x, y int) {
    px, py := ds.Find(x), ds.Find(y)
    if px == py {
        return
    }
    
    if ds.rank[px] < ds.rank[py] {
        ds.parent[px] = py
    } else if ds.rank[px] > ds.rank[py] {
        ds.parent[py] = px
    } else {
        ds.parent[py] = px
        ds.rank[px]++
    }
}
```

## Skip List

### Skip List Node
```go
type SkipListNode struct {
    value    int
    forward  []*SkipListNode
}

type SkipList struct {
    head     *SkipListNode
    level    int
    maxLevel int
}

func NewSkipList(maxLevel int) *SkipList {
    return &SkipList{
        head: &SkipListNode{
            forward: make([]*SkipListNode, maxLevel),
        },
        maxLevel: maxLevel,
    }
}
```

### Skip List Operations
```go
// Insert element
func (sl *SkipList) Insert(value int) {
    update := make([]*SkipListNode, sl.maxLevel)
    current := sl.head

    // Find position to insert
    for i := sl.level - 1; i >= 0; i-- {
        for current.forward[i] != nil && 
            current.forward[i].value < value {
            current = current.forward[i]
        }
        update[i] = current
    }

    // Generate random level
    level := sl.randomLevel()
    if level > sl.level {
        for i := sl.level; i < level; i++ {
            update[i] = sl.head
        }
        sl.level = level
    }

    // Create new node
    newNode := &SkipListNode{
        value:   value,
        forward: make([]*SkipListNode, level),
    }

    // Insert node at each level
    for i := 0; i < level; i++ {
        newNode.forward[i] = update[i].forward[i]
        update[i].forward[i] = newNode
    }
}
```

## Bloom Filter

### Bloom Filter Implementation
```go
type BloomFilter struct {
    bits    []bool
    k       int // Number of hash functions
    m       int // Size of bit array
}

func NewBloomFilter(size, k int) *BloomFilter {
    return &BloomFilter{
        bits: make([]bool, size),
        k:    k,
        m:    size,
    }
}

// Add element
func (bf *BloomFilter) Add(item string) {
    for i := 0; i < bf.k; i++ {
        position := bf.hash(item, i)
        bf.bits[position] = true
    }
}

// Check if element might exist
func (bf *BloomFilter) MightContain(item string) bool {
    for i := 0; i < bf.k; i++ {
        position := bf.hash(item, i)
        if !bf.bits[position] {
            return false
        }
    }
    return true
}
```

## LRU Cache

### LRU Implementation
```go
type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node
    tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     &Node{},
        tail:     &Node{},
    }
}
