package main

import (
	"bytes"
	"log"
	"net"
	"time"

	proxy "./proxy"
)

const (
	readTimeout = 10 * time.Second
	maxNBuf     = 2048
)

func setReadTimeout(c net.Conn) {
	if readTimeout != 0 {
		c.SetReadDeadline(time.Now().Add(readTimeout))
	}
}

func readData(conn net.Conn) *bytes.Buffer {
	result := new(bytes.Buffer)
	buf := make([]byte, maxNBuf)
	for {
		setReadTimeout(conn)
		n, err := conn.Read(buf)
		log.Println("PIPE DATA: ", n, err)
		if n > 0 {
			result.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return result
}

func main() {
	remote, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		log.Println("[Dial ERR]", err)
		return
	}
	defer remote.Close()

	getData := proxy.GetData("baidu.com")
	if _, err := remote.Write([]byte(getData)); err != nil {
		log.Println("[PIPE DATA ERROR]", err)
	}

	bytes := readData(remote)
	log.Println(bytes.String())
}
