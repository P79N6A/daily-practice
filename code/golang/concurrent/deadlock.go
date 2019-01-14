package main

import (
	"fmt"
)
func LogFmt(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func Log(args ...interface{}) {
	fmt.Println(args...)
}


func main() {
	a := make(chan int)
	go func() {
		a <- 1
	}()

	for i:= range a {
		Log(i)
	}
}
