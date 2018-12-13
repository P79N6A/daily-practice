package main

import (
	. ".."
)

func main() {
	n1 := NewNode(2, NewNode(3, NewNode(4, NewNode(5, nil))))
	n2 := NewNode(8, NewNode(7, NewNode(6, NewNode(5, nil))))

	n3 := n1.Reverse(nil)
	n4 := n2.Reverse(nil)
	n3.Drain()
	n4.Drain()

	Sum(n3, n4, 0).Reverse(nil).Drain()
}
