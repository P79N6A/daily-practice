package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	value := l1.Val + l2.Val
	result := &ListNode{value % 10, addTwoNumbers(l1.Next, l2.Next)}
	if value >= 10 {
		result.Next = addTwoNumbers(&ListNode{value / 10, nil}, result.Next)
	}
	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}
