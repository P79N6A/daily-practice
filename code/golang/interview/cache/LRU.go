// 双向链表+散列表，实现 LRU cache
package cache

import "fmt"

type Value interface{}

type biNode struct {
	prev  *biNode
	next  *biNode
	key   string
	value Value
}

func newNode(k string, v Value, prev *biNode, next *biNode) *biNode {
	return &biNode{
		key:   k,
		value: v,
		prev:  prev,
		next:  next,
	}
}

type LRUCache struct {
	data     map[string]*biNode
	capacity int
	head     *biNode
	tail     *biNode
}

func (lru *LRUCache) move2tail(node *biNode) {
	lru.tail.prev.next = node
	lru.tail.prev = node
	node.prev = lru.tail.prev
	node.next = lru.tail
}

func (lru *LRUCache) Get(k string) (Value, bool) {
	node, ok := lru.data[k]
	if !ok {
		return "", ok
	}
	node.prev.next = node.next
	node.next.prev = node.prev

	lru.move2tail(node)

	return node.value, true
}

func (lru *LRUCache) Set(k string, v Value) {
	size := len(lru.data)
	if size == lru.capacity {
		delete(lru.data, lru.head.next.key)
		lru.head.next = lru.head.next.next
		lru.head.next.prev = lru.head
	}
	node := newNode(k, v, nil, nil)
	lru.data[k] = node
	lru.move2tail(node)
}

func (lru *LRUCache) Print() {
	for k, v := range lru.data {
		fmt.Println(k, v)
	}
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		data:     make(map[string]*biNode),
	}
	lru.head = newNode("", "", nil, lru.tail)
	lru.tail = newNode("", "", lru.head, nil)
	return lru
}
