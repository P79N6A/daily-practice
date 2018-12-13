package main

import (
	"log"
	"net"
	"os"
	"time"
)

var (
	n       int
	timeout = 10 * time.Second
)

func say(conn net.Conn, words string) (n int, err error) {
	n, err = conn.Write([]byte(words))
	return
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9998")
	if err != nil {
		log.Println("[ERR]", err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept Error: %v\n", err)
			return
		}
		if err != nil {
			log.Printf("SetKeepAlive Error: %v\n", err)
		}
		n, err = say(conn, "hello")
		if err != nil {
			log.Printf("Write Error: %v\n", err)
		}
		log.Println("[WRITE N]", n)
		conn.SetWriteDeadline(time.Now().Add(timeout))
		time.Sleep(3 * time.Second)
		n, err = say(conn, "hello")
		log.Println("[WRITE N]", n)
		if err != nil {
			log.Printf("Write Error: %v\n", err)
		}
	}
}
