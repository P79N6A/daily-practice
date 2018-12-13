package main

import (
    "fmt"
    "errors"
)

func errorProne(url string) (name string, err error) {
    name = "holy"
    err = errors.New("crap")
    return
}

func main() {
    name, err := errorProne("http://holy.crap")
    fmt.Println("name is:", name, "err is", err)
}
