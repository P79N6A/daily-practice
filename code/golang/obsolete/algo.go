package main

import (
    "fmt"
)

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func fibonacci(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}

func main() {
    x, y := 16, 8
    fmt.Printf("gcd(%d,%d)=%d\n", x, y, gcd(x, y))
    fmt.Printf("fibonacci(%d)=%d\n", y, fibonacci(y))
}
