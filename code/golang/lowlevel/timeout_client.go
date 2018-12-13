package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9998")
	if err != nil {
		log.Println("[CONN ERR]", err)
	}

	io.Copy(os.Stdout, conn)
	time.Sleep(30 * time.Second)
	defer conn.Close()
}
