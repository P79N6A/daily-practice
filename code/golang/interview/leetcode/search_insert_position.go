// https://leetcode.com/problems/search-insert-position/description/
package leetcode

func searchInsert(nums []int, target int) (pos int) {
	for i, v := range nums {
		if v == target {
			return i
		} else if target > v {
			pos = i + 1
		}
	}
	return
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
