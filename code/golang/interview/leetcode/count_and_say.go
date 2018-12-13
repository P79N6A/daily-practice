// https://leetcode.com/problems/count-and-say/description/
package leetcode

func countAndSay(n int) string {
	return gen("1", n-1)
}

func gen(str string, n int) string {
	if n == 0 {
		return str
	}

	var ret []byte
	idx := 0
	for idx < len(str) {
		c := str[idx]
		lastJ := idx
		for j := idx; j < len(str); j++ {
			if str[j] != c {
				break
			}
			lastJ = j
		}

		n := lastJ - idx + 1
		ret = append(ret, byte(n)+'0')
		ret = append(ret, c)
		idx = lastJ + 1
	}

	return gen(string(ret), n-1)
}

func CountAndSay(n int) string {
	return countAndSay(n)
}
