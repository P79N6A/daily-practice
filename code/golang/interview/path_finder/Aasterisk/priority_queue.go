package Aasterisk

import (
	"container/heap"
	"fmt"
)

type Value interface{}

type Item struct {
	Value    Value
	Priority int
}

func (item Item) String() string {
	return fmt.Sprintf("(Priority: %d, %v)", item.Priority, item.Value)
}

type priorityQueueImpl struct {
	data    []*Item
	compare func(i, j int) bool
}

func (pq priorityQueueImpl) Len() int { return len(pq.data) }

func (pq priorityQueueImpl) Less(i, j int) bool {
	return pq.compare(pq.data[i].Priority, pq.data[j].Priority)
}

func (pq priorityQueueImpl) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityQueueImpl) Push(x interface{}) {
	item := x.(*Item)
	pq.data = append(pq.data, item)
}

func (pq *priorityQueueImpl) Pop() interface{} {
	old := pq.data
	n := len(old)
	item := old[n-1]
	pq.data = old[0 : n-1]
	return item
}

func (pq *priorityQueueImpl) update(item *Item, value Value, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, len(pq.data))
}

type PriorityQueue struct {
	*priorityQueueImpl
}

func (p *PriorityQueue) Insert(item *Item) {
	heap.Push(p, item)
}

func (p *PriorityQueue) Remove() *Item {
	return heap.Pop(p).(*Item)
}

func (p *PriorityQueue) Peek() *Item {
	return p.data[0]
}

func NewItem(v Value, p int) *Item {
	return &Item{
		v,
		p,
	}
}

func NewPriorityQueue(item *Item, compare func(i, j int) bool) *PriorityQueue {
	pq := priorityQueueImpl{data: make([]*Item, 0)}
	pq.compare = compare
	heap.Init(&pq)
	heap.Push(&pq, item)
	return &PriorityQueue{&pq}
}
