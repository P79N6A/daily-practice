package main

import (
    "fmt"
)

func foo(args ...interface{}) {
    for idx, v := range args {
        fmt.Println(idx, v)
    }
}

func main() {
    foo(1, 2, "str", true)
}
