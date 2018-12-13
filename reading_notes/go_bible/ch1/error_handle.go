package main

import (
    "fmt"
)

func holy(args ...interface{}){
    for _, v := range args {
        if v == "hallelujah" {
            panic("[holy] crusader: amen")
        }
    }
}
func foo(args ...interface{}) {
    holy(args...)
}
func fight(args ...interface{}) {
    foo(args...)
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("may be Gospel:\n\t %v\n", r)
        }
    } ()

    fight(1, true, "hallelujah")
}
