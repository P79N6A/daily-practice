package main

import (
    "fmt"
)

var stack []string

func push(elem string) {
    stack = append(stack, elem)
}

func pop() string {
    elem := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    return elem
}

func main() {
    push("hello")
    push("world")
    push("buddy")
    push("programming")
    last := pop()
    fmt.Println("last:", last)
    fmt.Println(stack)
}
