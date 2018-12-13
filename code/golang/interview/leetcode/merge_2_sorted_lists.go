package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	smaller := l1
	bigger := l2
	if l1.Val > l2.Val {
		smaller = l2
		bigger = l1
	}
	return &ListNode{smaller.Val, mergeTwoLists(smaller.Next, bigger)}
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}
