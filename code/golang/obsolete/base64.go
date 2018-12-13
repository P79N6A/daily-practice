package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

var encode = flag.String("encode", "", "encode string to base64")
var decode = flag.String("decode", "", "decode string from base64")

type String struct {
	Value string
}

func (s String) length() int {
	return len(s.Value)
}
func (s String) ToBase64() (string, error) {
	data := []byte(s.Value)
	return base64.StdEncoding.EncodeToString(data), nil
}
func (s String) FromBase64() (string, error) {
	data, err := base64.StdEncoding.DecodeString(s.Value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

func main() {
	flag.Parse()

	encoded, decoded := String{*encode}, String{*decode}

	if encoded.length() > 0 {
		str, err := encoded.ToBase64()
		if err != nil {
			printError(err)
			return
		}
		fmt.Printf("%s\n", str)
	} else if decoded.length() > 0 {
		str, err := decoded.FromBase64()
		if err != nil {
			printError(err)
			return
		}
		fmt.Printf("%s\n", str)
	}
}
