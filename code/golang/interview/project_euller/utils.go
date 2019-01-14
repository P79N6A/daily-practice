package project_euller

import "fmt"

func LogFmt(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Log(args ...interface{}) {
	fmt.Println(args...)
}

func Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func SumSlice(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}