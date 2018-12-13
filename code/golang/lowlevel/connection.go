package main

import (
	"log"
	"net"
	"os"
	"time"
)

var (
	network = "tcp"
	host    = "example.org:80"
	one     = make([]byte, 0)
)

func IsClosed(c net.Conn) (result bool) {
	if c == nil {
		return true
	}
	c.SetReadDeadline(time.Now())
	if _, err := c.Read(one); err != nil {
		log.Print("[CONN ERROR]", err)
		c.Close()
		c = nil
		result = true
	} else {
		c.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	}
	return
}

func main() {
	conn, err := net.Dial(network, host)
	if err != nil {
		log.Println("[Error]", err)
		os.Exit(1)
	}

	log.Println("[CONN]", conn)

	time.Sleep(1 * time.Second)

	conn.Close()

	time.Sleep(1 * time.Second)

	isClosed := IsClosed(conn)
	log.Println("[CONN after CLOSE]", isClosed, conn == nil)
}
