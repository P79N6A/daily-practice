package main

import (
	"log"
)

func deadlock() {
	ch := make(chan int)

	ch <- 2
	x := <-ch
	log.Println(x)
}

func nolock() {
	ch := make(chan int)

	go func() {
		ch <- 2
		log.Println("after write")
	}()

	x := <-ch
	log.Println("after read:", x)
}

func main() {
	nolock()
}
