// Package main demonstrates advanced data structures in Go
package main

import (
	"container/list"
	"errors"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
)

// MinHeap implements a min heap data structure
type MinHeap struct {
	items []int
	size  int
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: make([]int, 0),
	}
}

// Insert adds an item to the heap
func (h *MinHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.size++
	h.heapifyUp(h.size - 1)
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() (int, error) {
	if h.size == 0 {
		return 0, errors.New("heap is empty")
	}

	min := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.items = h.items[:h.size]
	if h.size > 0 {
		h.heapifyDown(0)
	}

	return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[parent] > h.items[index] {
			h.items[parent], h.items[index] = h.items[index], h.items[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) heapifyDown(index int) {
	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < h.size && h.items[left] < h.items[smallest] {
			smallest = left
		}
		if right < h.size && h.items[right] < h.items[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}

// Trie implements a trie data structure
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

// NewTrie creates a new trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word to the trie
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

// Search checks if a word exists in the trie
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

// StartsWith checks if any word starts with the given prefix
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

// DisjointSet implements a union-find data structure
type DisjointSet struct {
	parent []int
	rank   []int
}

// NewDisjointSet creates a new disjoint set
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

// Find returns the representative of the set containing x
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// Union combines the sets containing x and y
func (ds *DisjointSet) Union(x, y int) {
	px, py := ds.Find(x), ds.Find(y)
	if px == py {
		return
	}

	// Union by rank
	if ds.rank[px] < ds.rank[py] {
		ds.parent[px] = py
	} else if ds.rank[px] > ds.rank[py] {
		ds.parent[py] = px
	} else {
		ds.parent[py] = px
		ds.rank[px]++
	}
}

// BloomFilter implements a bloom filter
type BloomFilter struct {
	bits []bool
	k    int // Number of hash functions
}

// NewBloomFilter creates a new bloom filter
func NewBloomFilter(size, k int) *BloomFilter {
	return &BloomFilter{
		bits: make([]bool, size),
		k:    k,
	}
}

// Add adds an item to the bloom filter
func (bf *BloomFilter) Add(item string) {
	for i := 0; i < bf.k; i++ {
		position := bf.hash(item, i) % len(bf.bits)
		bf.bits[position] = true
	}
}

// Contains checks if an item might be in the set
func (bf *BloomFilter) Contains(item string) bool {
	for i := 0; i < bf.k; i++ {
		position := bf.hash(item, i) % len(bf.bits)
		if !bf.bits[position] {
			return false
		}
	}
	return true
}

// hash generates different hash values for the same input
func (bf *BloomFilter) hash(s string, seed int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	h.Write([]byte{byte(seed)})
	return h.Sum32()
}

// LRUCache implements a Least Recently Used cache
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	lru      *list.List
}

type entry struct {
	key   string
	value interface{}
}

// NewLRUCache creates a new LRU cache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		lru:      list.New(),
	}
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key string) (interface{}, bool) {
	if elem, ok := c.cache[key]; ok {
		c.lru.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// Put adds a value to the cache
func (c *LRUCache) Put(key string, value interface{}) {
	if elem, ok := c.cache[key]; ok {
		c.lru.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if c.lru.Len() >= c.capacity {
		oldest := c.lru.Back()
		if oldest != nil {
			delete(c.cache, oldest.Value.(*entry).key)
			c.lru.Remove(oldest)
		}
	}

	elem := c.lru.PushFront(&entry{key, value})
	c.cache[key] = elem
}

func main() {
	// Demonstrate MinHeap
	heap := NewMinHeap()
	fmt.Println("Adding numbers to MinHeap: 5, 3, 7, 1, 4")
	numbers := []int{5, 3, 7, 1, 4}
	for _, n := range numbers {
		heap.Insert(n)
	}
