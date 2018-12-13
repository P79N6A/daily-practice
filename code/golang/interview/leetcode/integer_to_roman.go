// https://leetcode.com/problems/integer-to-roman/description/
package leetcode

import "strings"

func intToRoman(num int) (repr string) {
	dict := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	ranks := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, rank := range ranks {
		qty := num / rank
		repr += strings.Repeat(dict[rank], qty)
		num = num % rank
	}

	return
}

func IntToRoman(num int) string {
	return intToRoman(num)
}
