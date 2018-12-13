// https://leetcode.com/problems/reverse-integer/description/
package leetcode

import "strconv"

func reverseStr(s string) (r string) {
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return
}

func reverse(x int) int {
	s := strconv.Itoa(abs(x))
	rs := reverseStr(s)
	if x < 0 {
		rs = "-" + rs
	}
	r, err := strconv.ParseInt(rs, 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}

func Reverse(x int) int {
	return reverse(x)
}
