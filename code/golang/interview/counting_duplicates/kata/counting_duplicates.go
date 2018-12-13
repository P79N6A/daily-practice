// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go
package kata

import "strings"

func duplicate_count(s1 string) (result int) {
	counts := make(map[rune]uint)

	for _, k := range strings.ToLower(s1) {
		if counts[k]++; counts[k] == 2 {
			result++
		}
	}
	return
}
