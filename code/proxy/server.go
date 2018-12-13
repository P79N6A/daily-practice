package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	proxy "./proxy"
)

const (
	network = "tcp"
	laddr   = ":10086"

	addrMask byte = 0xf

	layoutType = 0
	layoutLen  = 1
	layoutIP   = 1
	layoutAddr = 2

	typeIPv4   = 1 // type is ipv4 address
	typeDomain = 3 // type is domain address
	typeIPv6   = 4 // type is ipv6 address

	lenIPv4     = net.IPv4len + 2 // ipv4 + 2port
	lenIPv6     = net.IPv6len + 2 // ipv6 + 2port
	lenPort     = 2
	lenHashMac1 = 10
	requestBuf  = 269
	// httpGET = "GET / HTTP/1.1\r\nHost: example.org\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36\r\n\r\n"
)

// request layout:
// 1(addrType) + 1(lenByte) + 255(max length address) + 2(port) + 10(hmac-sha1)
func extractRequest(conn net.Conn) (host string, err error) {
	buf := make([]byte, requestBuf)

	if _, err = io.ReadFull(conn, buf[:layoutType+1]); err != nil {
		return
	}

	var reqStart, reqEnd int
	addrType := buf[layoutType]
	switch addrType & addrMask {
	case typeIPv4:
		reqStart, reqEnd = layoutIP, layoutIP+lenIPv4
	case typeIPv6:
		reqStart, reqEnd = layoutIP, layoutIP+lenIPv6
	case typeDomain:
		if _, err = io.ReadFull(conn, buf[layoutLen:layoutLen+1]); err != nil {
			return
		}
		domainLen := int(buf[layoutLen])
		reqStart, reqEnd = layoutAddr, layoutAddr+domainLen+lenPort
	default:
		err = fmt.Errorf("unsupported addr type:%d", addrType&addrMask)
	}

	if _, err = io.ReadFull(conn, buf[reqStart:reqEnd]); err != nil {
		return
	}

	switch addrType & addrMask {
	case typeIPv4:
		host = net.IP(buf[reqStart:reqEnd]).String()
	case typeIPv6:
		host = net.IP(buf[reqStart:reqEnd]).String()
	case typeDomain:
		host = string(buf[reqStart : reqEnd-2])
	}

	port := binary.BigEndian.Uint16(buf[reqEnd-2 : reqEnd])
	host = net.JoinHostPort(host, strconv.Itoa(int(port)))

	return
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1. act as a transparent proxy
	host, err := extractRequest(conn)
	if err != nil {
		log.Println("[PARSE REQ ERROR]", err)
		return
	}

	remote, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("[DIAL REMOTE ERROR]", err)
		return
	}

	defer remote.Close()

	log.Println("[REQUEST TO]", host)

	// 2. act as a tunnel
	go proxy.Pipe(conn, remote)
	proxy.Pipe(remote, conn)
	return
}

func main() {
	log.Println("server configuration:", network, laddr)
	ln, err := net.Listen(network, laddr)
	if err != nil {
		log.Println("error listening")
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			return
		}

		log.Println("[CONN COMING]!!!")
		handleConn(conn)
	}
}
