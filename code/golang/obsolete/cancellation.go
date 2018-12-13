package main

import (
	"log"
	"os"
)

var done = make(chan struct{})

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for {
		select {
		case i := <-done:
			log.Println("RECEIVED: ", i)
			return
		}
	}
}
