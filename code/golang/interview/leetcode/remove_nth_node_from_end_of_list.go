// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)
	if headIsNthFromEnd {
		return head.Next
	}
	d.Next = d.Next.Next

	return head
}

func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head

	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}
		n--
		head = head.Next
	}
	headIsNthFromEnd = n == 0

	return
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}