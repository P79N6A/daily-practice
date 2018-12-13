// https://leetcode.com/problems/container-with-most-water/description/
package leetcode

func maxArea(height []int) (max int) {
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			m := min(height[i], height[j])
			area := m * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return
}

func MaxArea(height []int) int {
	return maxArea(height)
}
