// https://www.ietf.org/rfc/rfc1928.txt

package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	proxy "./proxy"
)

const (
	network    = "tcp"
	laddr      = "127.0.0.1:8088"
	requestBuf = 269

	socksVersion     = 5
	layoutVer        = 0
	layoutNofMethods = 1
	layoutMethods    = 2
	layoutCommand    = 1
	layoutRSV        = 2
	layoutATYP       = 3
	layoutAddr       = 4
	CONNECT          = 1

	typeIPv4   = uint8(1)        // type is ipv4 address
	typeDomain = uint8(3)        // type is domain address
	typeIPv6   = uint8(4)        // type is ipv6 address
	lenIPv4    = net.IPv4len + 2 // ipv4 + 2port
	lenIPv6    = net.IPv6len + 2 // ipv6 + 2port
)

func reply(conn net.Conn, bytes []byte) (err error) {
	if _, err := conn.Write(bytes); err != nil {
		log.Println("[RESP Error]", err)
	}
	return
}

// extract method negotiation header
// |VER | NMETHODS | METHODS  |
// +----+----------+----------+
// | 1  |    1     | 1 to 255 |
func extractNeogotiation(conn net.Conn) (socks *proxy.Socks5Negotiation, err error) {
	buf := make([]byte, requestBuf)

	if _, err = io.ReadFull(conn, buf[:layoutNofMethods+1]); err != nil {
		return
	}

	version := uint8(buf[layoutVer])
	if version != 5 {
		err = errors.New("NOT a Socks5 request")
		return
	}
	nOfMethods := uint8(buf[layoutNofMethods])
	if nOfMethods == 0 {
		// do nothing
	} else if _, err = io.ReadFull(conn, buf[layoutMethods:layoutMethods+nOfMethods]); err != nil {
		return
	}

	socks = new(proxy.Socks5Negotiation)
	socks.Version = version
	socks.NumOfMethods = nOfMethods

	for i := uint8(0); i < nOfMethods; i++ {
		method := uint8(buf[layoutMethods+i])
		socks.Methods = append(socks.Methods, method)
	}

	return
}

// extract socks5 request
// +----+-----+-------+------+----------+----------+
// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
// +----+-----+-------+------+----------+----------+
func extractRequest(conn net.Conn) (socksReq *proxy.Socks5Request, err error) {
	buf := make([]byte, requestBuf)

	if _, err = io.ReadFull(conn, buf[:layoutATYP+1]); err != nil {
		return
	}

	version := uint8(buf[layoutVer])
	if version != 5 {
		err = errors.New("NOT a socks5 request")
		return
	}
	command := uint8(buf[layoutCommand])
	if command != CONNECT {
		err = errors.New("only CONNECT be able to accept!")
		return
	}
	RSV := uint8(buf[layoutRSV])
	ATYP := uint8(buf[layoutATYP])
	var address string
	var addrEnd int

	switch ATYP {
	case typeIPv4:
		if _, err = io.ReadFull(conn, buf[layoutAddr:layoutAddr+lenIPv4]); err != nil {
			return
		}
		addrEnd = layoutAddr + lenIPv4
		address = net.IP(buf[layoutAddr:addrEnd]).String()
	case typeIPv6:
		if _, err = io.ReadFull(conn, buf[layoutAddr:layoutAddr+lenIPv6]); err != nil {
			return
		}
		addrEnd = layoutAddr + lenIPv6
		address = net.IP(buf[layoutAddr:addrEnd]).String()
	case typeDomain:
		if _, err = io.ReadFull(conn, buf[layoutAddr:layoutAddr+1]); err != nil {
			return
		}
		addrLen := int(buf[layoutAddr])
		addrEnd = layoutAddr + 1 + addrLen
		if _, err = io.ReadFull(conn, buf[layoutAddr+1:addrEnd]); err != nil {
			return
		}
		address = string(buf[layoutAddr+1 : addrEnd])
	default:
		err = fmt.Errorf("address type is Unknown: %d", ATYP)
		return
	}

	if _, err = io.ReadFull(conn, buf[addrEnd:addrEnd+2]); err != nil {
		return
	}

	socksReq = new(proxy.Socks5Request)
	socksReq.Version = version
	socksReq.Command = command
	socksReq.RSV = RSV
	socksReq.AddressType = ATYP
	socksReq.Address = address
	socksReq.Port = binary.BigEndian.Uint16(buf[addrEnd : addrEnd+2])
	socksReq.AddressWithPort = net.JoinHostPort(address, strconv.Itoa(int(socksReq.Port)))
	return
}

// reply the method selection
// |VER | METHOD |
// +----+--------+
// | 1  |   1    |
func replyNegotiation(conn net.Conn, socks *proxy.Socks5Negotiation) {
	// no authentication required
	reply(conn, []byte{socksVersion, 0x00})
	// select the first method, always
	// reply(conn, []byte{socksVersion, socks.Methods[0]})
}

// reply the request
// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
func replyRequest(conn net.Conn, socksRequest *proxy.Socks5Request) {
	reply(conn, []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x80, 0x88})
}

// rountine of per connection
func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1. extract info about negotiation
	socks, err := extractNeogotiation(conn)
	if err != nil {
		log.Println("[Negotiate REQ Error]", err)
		return
	}
	log.Println(socks)

	// 2. confirm negotiation
	replyNegotiation(conn, socks)

	// 3. extract info about request
	socksRequest, err := extractRequest(conn)
	if err != nil {
		log.Println("[Extract REQ Error]", err)
		return
	}
	log.Println(socksRequest)

	// 4. confirm the connection was established
	replyRequest(conn, socksRequest)

	// 5. connect to remote
	if socksRequest.Port == 80 {
		go func() { // exhaust stream of original request
			_ = proxy.ReadStream(conn)
		}()

		// request to remote
		domain := "baidu.com"
		remote, err := proxy.ConnectWithHTTP(domain, 80)
		if err != nil {
			log.Println("CONNECT ERR WITH", domain, err)
			return
		}

		// respond to client
		proxy.Pipe(remote, conn)
	}

	return
}

func main() {
	log.Println("Server is Listening:", network, laddr)
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

		log.Println("[CONN]")
		handleConn(conn)
	}
}
