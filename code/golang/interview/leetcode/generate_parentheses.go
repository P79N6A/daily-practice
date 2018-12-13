// https://leetcode.com/problems/generate-parentheses
package leetcode

func addParenthesis(l, r int, out string, parentheses *[]string) {
	if l < 0 || r < 0 || l > r {
		return
	}
	if l == 0 && r == 0 {
		*parentheses = append(*parentheses, out)
		return
	}
	addParenthesis(l-1, r, out+"(", parentheses)
	addParenthesis(l, r-1, out+")", parentheses)
}
func generateParenthesis(n int) (parentheses []string) {
	addParenthesis(n, n, "", &parentheses)
	return
}

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}
