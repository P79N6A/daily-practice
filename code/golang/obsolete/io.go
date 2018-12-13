package main

import (
    "fmt"
    "flag"
    "os"
)

var filename = flag.String("f", "/etc/hosts", "input a filename")

func main() {
    flag.Parse()
    f, err := os.Open(*filename)
    if err != nil {
        return
    }
    fstat, _ := f.Stat()
    fmt.Println(fstat)
    f.Close()

    cwd, _ := os.Getwd()
    fmt.Println(cwd)
}
