package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"log"
	"net"
	"os"
	"strings"
)

const (
	network = "tcp"
	laddr   = ":10086"

	port uint16 = 80

	typeDomain = 3
	maxNBuf    = 2048
)

func sendProtocolHeader(conn net.Conn, domain string) {
	buf := new(bytes.Buffer)

	buf.WriteByte(typeDomain)
	buf.WriteByte(byte(len(domain)))
	buf.WriteString(domain)
	binary.Write(buf, binary.BigEndian, port)
	request := buf.Bytes()
	log.Printf("[REQUEST IN BigEndian] %v\n", request)

	binary.Write(conn, binary.BigEndian, &request)
}

func sendProtocolContent(conn net.Conn, contents []byte) {
	if _, err := conn.Write(contents); err != nil {
		log.Println("[HTTP GET ERROR]", err)
		return
	}
}

func httpContent(domain string) string {
	return strings.Join([]string{"GET / HTTP/1.1\r\n",
		"Host: " + domain + "\r\n",
		"Connection: Close\r\n",
		"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6)\r\n\r\n"}, "")
}

func handleResponse(conn net.Conn, domain string) {
	var total bytes.Buffer
	buf := make([]byte, maxNBuf)
	log.Println("[REQUEST Via Proxy]", domain)

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			total.Write(buf)
		}
		if err != nil {
			break
		}
	}

	log.Println("\n\n", total.String())
}

// request layout:
// 1(addrType) + 1(lenByte) + 255(max length address) + 2(port) + 10(hmac-sha1)
func main() {
	domain := flag.String("domain", "example.org", "domain u want to visit")

	conn, err := net.Dial("tcp", laddr)
	if err != nil {
		log.Println("[ERROR]", err)
		os.Exit(1)
	}

	defer conn.Close()

	sendProtocolHeader(conn, *domain)

	sendProtocolContent(conn, []byte(httpContent(*domain)))

	handleResponse(conn, *domain)
}
