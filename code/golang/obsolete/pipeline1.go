package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

// catch panic
func catchPanic() {
	if p := recover(); p != nil {
		log.Printf("caught panic: %v\n%s\n", p, debug.Stack())
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		defer catchPanic()
		for x := 1; ; x++ {
			naturals <- x
		}
	}()
	go func() {
		for {
			for x := range naturals {
				squares <- x * x
			}
			close(squares)
			break
		}
	}()
	for {
		for x := range squares {
			if x == 65536 {
				close(naturals)
			}
			fmt.Println(x)
			time.Sleep(800 * time.Microsecond)
		}
		break
	}
}
