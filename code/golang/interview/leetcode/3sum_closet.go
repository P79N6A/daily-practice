// https://leetcode.com/problems/3sum-closest/description/
package leetcode

func threeSumClosest(nums []int, target int) int {
	approximation := 99999999

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum3 := nums[i] + nums[j] + nums[k]
				distance := abs(sum3 - target)
				if distance < abs(approximation-target) {
					approximation = sum3
				}
			}
		}
	}

	return approximation
}

func ThreeSumClosest(nums []int, target int) int {
	return threeSumClosest(nums, target)
}
