package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()

	client := make(chan string)
	last := time.Now()
	tick := time.NewTicker(1 * time.Second)

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			text := input.Text()
			client <- text
		}
		log.Println("CLOSE CHAN AFTER CLIENT WAS CLOSED")
		close(client)
	}()

	for {
		select {
		case <-tick.C:
			if time.Now().After(last.Add(10 * time.Second)) {
				tick.Stop()
				log.Println("CLIENT END")
				return
			}
		case text, ok := <-client:
			if ok {
				log.Print("CLIENT SAID: ", text)
				last = time.Now()
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}
