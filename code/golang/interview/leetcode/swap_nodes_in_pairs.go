// https://leetcode.com/problems/swap-nodes-in-pairs/description/
package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		next := head.Next
		nn := next.Next

		next.Next = head
		head = next

		head.Next.Next = swapPairs(nn)
		return head
	}
}

func SwapPairs(head *ListNode) *ListNode {
	return swapPairs(head)
}
