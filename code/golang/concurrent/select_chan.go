package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		seconds := 5 * time.Second
		sub := make(chan struct{})

		go func() {
			log.Printf("SLEEP %v\n", seconds)
			time.Sleep(seconds)
			log.Println("SUB<-")
			sub <- struct{}{}
		}()

		// emit a default value for channel
		defer close(ch)

		for {
			log.Println("LOOP AND SELECT")
			select {
			case <-sub:
				log.Println("<-SUB")
				return
			}
		}
	}()

	log.Println("WAIT FOR AWAKE...")
	log.Println("RECEIVE: ", <-ch)

	log.Println("END")
}
