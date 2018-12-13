// https://leetcode.com/problems/valid-parentheses/description/
package leetcode

type Stack struct {
	data []rune
}

func (s *Stack) push(ch rune) {
	s.data = append(s.data, ch)
}
func (s *Stack) len() int {
	return len(s.data)
}
func (s *Stack) peak() (last rune) {
	l := len(s.data)
	if l == 0 {
		return
	}
	last = s.data[l-1]
	return
}
func (s *Stack) pop() (last rune) {
	l := len(s.data)
	if l == 0 {
		return
	}
	last = s.data[l-1]
	s.data = s.data[0 : l-1]
	return
}

func opposite(a, b rune) bool {
	opposite := map[string]string{
		"(": ")",
		")": "(",
		"{": "}",
		"}": "{",
		"[": "]",
		"]": "[",
	}
	if v, ok := opposite[string(a)]; ok {
		if v == string(b) {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	stack := Stack{data: []rune{}}
	for _, ch := range s {
		if stack.len() > 0 {
			if opposite(stack.peak(), ch) {
				stack.pop()
				continue
			}
		}
		stack.push(ch)
	}
	return stack.len() == 0
}

func IsValid(s string) bool {
	return isValid(s)
}
