// https://leetcode.com/problems/trapping-rain-water/description/
package kata

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func trap(height []int) (water int) {
	l := make([]int, len(height))
	r := make([]int, len(height))
	for i, max := 0, 0; i < len(height); i++ {
		curr := height[i]
		if curr > max {
			max = curr
		}
		l[i] = max - curr
	}

	for j, max := len(height)-1, 0; j >= 0; j-- {
		curr := height[j]
		if curr > max {
			max = curr
		}
		r[j] = max - curr
	}

	for i := range height {
		water += min(l[i], r[i])
	}

	return
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

	fmt.Println(trap([]int{0, 1, 2, 1, 0, 2, 2, 0, 3, 2}))

	fmt.Println(trap([]int{0, 1, 0, 1, 0, 3, 0, 3, 0, 1}))
}
