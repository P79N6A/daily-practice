package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    for _, filename := range os.Args[1:] {
        f, err := os.Open(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            continue
        }
        input := bufio.NewScanner(f)
        for input.Scan() {
            fmt.Printf("%s\n", input.Text())
        }
    }
}
