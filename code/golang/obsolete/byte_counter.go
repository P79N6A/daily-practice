package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		fmt.Println("word:", x)
		words++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += LineCounter(words)
	return words, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println("chars:", c)

	var l LineCounter
	l.Write("Now is the winter,\nMade glorious summer of York.\n")
	fmt.Println("word:", l)
}
