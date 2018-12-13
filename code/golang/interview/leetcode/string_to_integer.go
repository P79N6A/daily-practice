// https://leetcode.com/problems/string-to-integer-atoi/description/
package leetcode

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	trimed := bytes.Trim([]byte(str), " ")
	if len(trimed) == 0 {
		return 0
	}
	first := rune(string(trimed)[0])
	if !(unicode.IsDigit(first) || first == '-' || first == '+') {
		return 0
	}
	reg := regexp.MustCompile(`([\+\-\sa-z]*[0-9]+)`)
	a := reg.FindString(string(trimed))
	i, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		if strings.Contains(err.Error(), "value out of range") {
			if i > 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		return 0
	}
	return int(i)
}

func MyAtoi(str string) int {
	return myAtoi(str)
}
