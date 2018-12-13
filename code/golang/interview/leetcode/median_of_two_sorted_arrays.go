// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2, len1, len2 := 0, 0, len(nums1), len(nums2)
	total := len1 + len2
	merged := make([]int, total)

	for i := 0; i < total; i++ {
		if l1 == len1 {
			merged[i] = nums2[l2]
			l2++
		} else if l2 == len2 {
			merged[i] = nums1[l1]
			l1++
		} else if nums1[l1] > nums2[l2] {
			merged[i] = nums2[l2]
			l2++
		} else if nums1[l1] <= nums2[l2] {
			merged[i] = nums1[l1]
			l1++
		}
	}

	if total%2 == 0 {
		return float64(merged[total/2]+merged[total/2-1]) / 2
	}
	return float64(merged[total/2])
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}
