// https://leetcode.com/problems/longest-common-prefix/description/
package leetcode

func prefix(a, b string) (prefix string) {
	min := len(a)
	if len(a) > len(b) {
		min = len(b)
	}
	for i := 0; i < min; i++ {
		if a[i] != b[i] {
			return
		}
		prefix += string(a[i])
	}
	return
}

func longestCommonPrefix(strs []string) (lcp string) {
	if len(strs) > 0 {
		lcp = strs[0]
		for i := 1; i < len(strs); i++ {
			lcp = prefix(lcp, strs[i])
		}
	}
	return
}

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}
