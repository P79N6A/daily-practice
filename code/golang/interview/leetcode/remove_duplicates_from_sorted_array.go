// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package leetcode

func removeDuplicates(nums []int) (count int) {
	uniqueness := make(map[int]bool)
	result := make([]int, 0)

	for _, n := range nums {
		if _, ok := uniqueness[n]; !ok {
			result = append(result, n)
			count++
			uniqueness[n] = true
		}
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	return
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}
