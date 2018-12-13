// https://leetcode.com/problems/merge-k-sorted-lists/description/
package leetcode

func partition(lists []*ListNode) (min *ListNode, rest []*ListNode) {
	for _, lst := range lists {
		if lst == nil {
			continue
		} else if min == nil || min.Val > lst.Val {
			min = lst
		}
	}
	for i, lst := range lists {
		if lst == nil {
			continue
		} else if lst.Val == min.Val {
			min = &ListNode{lst.Val, nil}
			lst = lst.Next
			if lst == nil {
				rest = append(lists[0:i], lists[i+1:]...)
			} else {
				lists[i] = lst
				rest = lists
			}
			break
		}
	}
	return
}

func mergeKLists(lists []*ListNode) *ListNode {
	min, rest := partition(lists)
	if len(rest) == 0 {
		return min
	} else {
		return &ListNode{min.Val, mergeKLists(rest)}
	}
}

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}
