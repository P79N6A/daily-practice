// https://leetcode.com/problems/two-sum/description/
package leetcode

func twoSum(nums []int, target int) (solution []int) {
	expected := make(map[int]int)

	for i, v := range nums {
		if j, ok := expected[v]; ok {
			solution = append(solution, j)
			solution = append(solution, i)
			return
		}
		expected[target-v] = i
	}

	return
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
