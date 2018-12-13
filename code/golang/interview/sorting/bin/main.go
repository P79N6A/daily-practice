package main

import (
	"fmt"

	. ".."
)

func main() {
	a := []int{11, 21, 3, 5, 9, 16, 23, 17}
	InsertionSort(a)
	fmt.Println(a)

	b := []int{11, 21, 3, 5, 9, 16, 23, 17}
	QuickSort(b)
	fmt.Println(b)
}
