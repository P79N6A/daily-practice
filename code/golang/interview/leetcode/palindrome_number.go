// https://leetcode.com/problems/palindrome-number/description/
package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	return isPalindrome(strconv.Itoa(x))
}
