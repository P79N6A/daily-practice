// https://leetcode.com/problems/roman-to-integer/description/
package leetcode

func romanToInt(s string) (total int) {
	dict := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	for i := 0; i < len(s); {
		end := i + 2
		if end > len(s) {
			end = i + 1
		}
		if n, ok := dict[s[i:end]]; ok {
			i = end
			total += n
		} else if n, ok := dict[s[i:end-1]]; ok {
			i = end - 1
			total += n
		}
	}

	return
}

func RomanToInt(s string) int {
	return romanToInt(s)
}
