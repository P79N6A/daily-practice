package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
)

func Bytes2Hex(bs []byte) string {
	hexes := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(hexes, bs)

	var buf bytes.Buffer
	buf.WriteString("0x ")
	if len(hexes)%2 != 0 {
		hexes = append([]byte{0}, hexes...)
	}
	for i := 0; i < len(hexes)/2; i++ {
		base := i * 2
		buf.WriteString(fmt.Sprintf("%c%c ", hexes[base], hexes[base+1]))
	}
	return buf.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		hexString := Bytes2Hex(input.Bytes())
		fmt.Println(hexString)
	}
}
