// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package leetcode

func letterCombinations(digits string) (combinations []string) {
	mapping := make(map[string][]string)
	mapping["2"] = []string{"a", "b", "c"}
	mapping["3"] = []string{"d", "e", "f"}
	mapping["4"] = []string{"g", "h", "i"}
	mapping["5"] = []string{"j", "k", "l"}
	mapping["6"] = []string{"m", "n", "o"}
	mapping["7"] = []string{"p", "q", "r", "s"}
	mapping["8"] = []string{"t", "u", "v"}
	mapping["9"] = []string{"w", "x", "y", "z"}

	dict := make([][]string, 0)
	total := 0

	for _, d := range digits {
		digit := string(d)
		if letters, ok := mapping[digit]; ok {
			dict = append(dict, letters)
			if total == 0 {
				total = len(letters)
			} else {
				total *= len(letters)
			}
		}
	}

	cursor := make([]int, len(dict))
	index := len(dict) - 1

	for x := 0; x < total; x++ {
		letters := ""
		for j := 0; j < len(dict); j++ {
			letters += dict[j][cursor[j]]
		}
		combinations = append(combinations, letters)

		cursor[index]++
		for i := index; i >= 0; i-- {
			if cursor[i] >= len(dict[i]) {
				cursor[i] = 0
				if i > 0 {
					cursor[i-1]++
				}
			}
		}
	}

	return
}

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}
