// https://leetcode.com/problems/remove-element/description/
package leetcode

func removeElement(nums []int, val int) int {
	removed := []int{}
	for _, v := range nums {
		if v != val {
			removed = append(removed, v)
		}
	}

	for i := 0; i < len(removed); i++ {
		nums[i] = removed[i]
	}

	return len(removed)
}

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
