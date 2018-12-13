package main

import (
	"fmt"
)

func main() {
	x := "hello"
	for _, x := range x {
		y := x + 'A' - 'a'
		fmt.Printf("%c", y) // "HELLO" (one letter per iteration)
	}
}
