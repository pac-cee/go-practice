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

