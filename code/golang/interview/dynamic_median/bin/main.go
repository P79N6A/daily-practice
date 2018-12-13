package main

import (
	. "../heap"
	"fmt"
)

func main() {
	array := []int{12, 100, 16, 4, 8, 70, 2, 36, 22, 5, 46, 56, 112, 233, 666}
	dm := Init(&array)

	fmt.Printf("中位数：%d\n", dm.Median())
	fmt.Println(dm.Drain())
}
