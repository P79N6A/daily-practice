// https://leetcode.com/problems/zigzag-conversion/description/
package leetcode

func convert(s string, numRows int) (result string) {
	if numRows == 1 {
		return s
	}
	matrix := make([]string, numRows)
	zigzagLen := 2*numRows - 2
	height := numRows - 1

	for i, j, round := 0, 0, 0; i < len(s); i++ {
		if j > zigzagLen {
			j = 0
		}
		if zigzagLen > 0 {
			round = i / zigzagLen
		}
		start := round * zigzagLen
		row := i - start

		if row > height {
			row = height - row%height
		} else {
			row = row % numRows
		}

		matrix[row] = matrix[row] + string(s[i])
		j++
	}

	for i := range matrix {
		result += matrix[i]
	}

	return
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}
