package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

type (
	TNTRequest struct {
		Type    uint8  // 0: invalid, 1: valid
		Length  uint16 // length of payload
		Payload []byte
	}
)

func (r *TNTRequest) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(r.Type)
	binary.Write(buf, binary.BigEndian, r.Length)
	buf.Write(r.Payload)
	return buf.Bytes()
}

func unmarshallRequest() (r *TNTRequest, err error) {
	r = new(TNTRequest)
	r.Type = uint8(2)
	return
}

func main() {
	const (
		domain           = "www.google.com.hk"
		domainLen        = len(domain)
		port      uint16 = 443
	)

	var buf = new(bytes.Buffer)
	buf.WriteByte(byte(domainLen))
	buf.WriteString(domain)
	binary.Write(buf, binary.BigEndian, port)

	payload := buf.Bytes()
	r := TNTRequest{1, uint16(len(payload)), payload}
	log.Println(r.Bytes())

	x, err := unmarshallRequest()
	log.Println(x, err)
}
