package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("PANIC: %v", err)
		}
	}()

	fmt.Println("RESULT: ", defer_call())
}
func defer_call() int {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("panicking")
	return 666
}
