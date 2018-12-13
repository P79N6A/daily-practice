package main

import (
	"log"
)

func main() {
	i := 1
	ch := make(chan int)
	go func() {
		defer func() {
			log.Println("GOROUTINE QUIT")
			ch <- 0
		}()

		for {
			log.Println("CHECK")
			if i < 10 {
				return
			}
		}
	}()

	<-ch
	log.Println("MAIN QUIT")
}
