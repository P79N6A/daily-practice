package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./algo.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		log.Print(line)
	}
}
