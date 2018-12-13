package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE.")
}
