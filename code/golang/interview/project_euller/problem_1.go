/*
Multiples of 3 and 5
Problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package project_euller

func Calc_1() {
	upper := 1000
	result := make([]int, 0)

	for i := 1; i < upper; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			result = append(result, i)
		}
	}

	Log(SumSlice(result))
}
