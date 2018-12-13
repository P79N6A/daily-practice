// https://leetcode.com/problems/implement-strstr/description/
package leetcode

func strStr(haystack string, needle string) (pos int) {
	pos = -1

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		same := 0
		for j := 0; j < len(needle); j++ {
			if j+i == len(haystack) {
				break
			}
			if haystack[j+i] == needle[j] {
				same++
			}
		}
		if same == len(needle) {
			return i
		}
	}

	return
}

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}
