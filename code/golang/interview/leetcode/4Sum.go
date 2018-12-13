// https://leetcode.com/problems/4sum/description/
package leetcode

func fourSum(nums []int, target int) (tetrads [][]int) {
	QuickSort(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum < target {
					k++
					for nums[k] == nums[k-1] && k < l {
						k++
					}
				} else if sum > target {
					l--
					for nums[l] == nums[l+1] && k < l {
						l--
					}
				} else {
					tetrads = append(tetrads, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for nums[k] == nums[k-1] && k < l {
						k++
					}
					for nums[l] == nums[l+1] && k < l {
						l--
					}
				}
			}
		}
	}

	return
}

func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
}
