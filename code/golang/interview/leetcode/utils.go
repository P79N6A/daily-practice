package leetcode

import "fmt"

func LogFmt(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Log(args ...interface{}) {
	fmt.Println(args...)
}

func Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1

	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Drain() {
	for n != nil {
		LogFmt("%d -> ", n.Val)
		n = n.Next
	}
	Log()
}

func revListNode(prev *ListNode, current *ListNode) *ListNode {
	if current.Next == nil {
		current.Next = prev
		return current
	} else {
		next := current.Next
		current.Next = prev
		return revListNode(current, next)
	}
}

func (n *ListNode) Reverse() *ListNode {
	return revListNode(nil, n)
}
