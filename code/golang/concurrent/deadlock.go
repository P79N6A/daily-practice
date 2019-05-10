package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	go func() {
		a <- 1
	}()

	for i := range a {
		fmt.Println(i)
	}
}
