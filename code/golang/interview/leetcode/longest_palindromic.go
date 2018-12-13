// https://leetcode.com/problems/longest-palindromic-substring/description/
package leetcode

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) (palindrome string) {
	longest := 0
	last := len(s) - 1

	for i := range s {
		for j := last; j >= i; j-- {
			word := s[i : j+1]
			if isPalindrome(word) && len(word) > longest {
				longest = len(word)
				palindrome = word
			}
		}
	}

	return
}

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}
