// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package leetcode

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) (longest int) {
	m := make([]rune, 256)
	left := 0

	for i, c := range s {
		if m[c] == 0 || int(m[c]) < left {
			longest = max(longest, i-left+1)
		} else {
			left = int(m[c])
		}
		m[c] = rune(i + 1)
	}

	return
}

func LengthOfLongestSubstring(s string) (longest int) {
	return lengthOfLongestSubstring(s)
}
