package leetcode

func gcd(a, b int) (res int) {
	if a < b {
		a, b = b, a
	}
	r := a % b
	if r == 0 {
		res = b
	} else {
		res = gcd(b, r)
	}
	return
}

func GCD(a, b int) (res int) {
	return gcd(a, b)
}
