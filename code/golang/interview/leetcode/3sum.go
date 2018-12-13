// https://leetcode.com/problems/3sum/description/
package leetcode

func threeSum(nums []int) (triplets [][]int) {
	QuickSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}
		}
	}

	return
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
