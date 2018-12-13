package main

import (
    "fmt"
)

func incr(p *int) int {
    *p++
    return *p
}

func newInt() *int {
    return new(int)
}

func newInt2() *int {
    var dummy int
    return &dummy
}

func main() {
    v := 1
    incr(&v)
    fmt.Println(incr(&v))

    fmt.Println(newInt() == newInt2())
}
