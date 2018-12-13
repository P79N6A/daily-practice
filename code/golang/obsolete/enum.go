package main

import (
    "fmt"
)

const (
    First = 1 << iota
    Second
    Third
    Fourth
    Fifth
    Sixth
    Seventh
    Eighth
)

func main() {
    fmt.Println(First, Second, Third, Fourth, Fifth, Sixth, Seventh, Eighth)
}
