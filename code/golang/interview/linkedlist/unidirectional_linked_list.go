package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) add(a *Node) (sum int, carried int) {
	sum = (n.Value + a.Value) % 10
	carried = (n.Value + a.Value) / 10
	return
}

func (n *Node) Drain() {
	for n != nil {
		fmt.Printf("%d -> ", n.Value)
		n = n.Next
	}
	fmt.Println()
}

func (n *Node) Reverse(prev *Node) *Node {
	if n.Next == nil {
		n.Next = prev
		return n
	}
	t := n.Next
	n.Next = prev
	return t.Reverse(n)
}

func Sum(a *Node, b *Node, c2 int) *Node {
	sum, carried := 0, 0
	if a != nil && b != nil {
		sum, carried = a.add(b)
		sum += c2
		return NewNode(sum, Sum(a.Next, b.Next, carried))
	} else if a != nil {
		sum, c2 = a.Value, 0
		return NewNode(sum, Sum(a.Next, nil, c2))
	} else if b != nil {
		sum, c2 = b.Value, 0
		return NewNode(sum, Sum(nil, b.Next, c2))
	} else if c2 > 0 {
		return NewNode(c2, nil)
	} else {
		return nil
	}
}

func NewNode(v int, next *Node) *Node {
	return &Node{
		Value: v,
		Next:  next,
	}
}
