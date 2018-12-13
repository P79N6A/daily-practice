package main

import (
    "fmt"
    "time"
)

type CalcResult struct {
    i int
    x int
}

func main() {
    ch := make(chan CalcResult)
    timeout := make(chan bool, 1)
    upper := 10
    for i:=0; i < upper; i++ {
        go pow(i, ch)
        go timer(timeout)
    }

    for i:=0; i < upper; i++ {
        select {
        case res := <-ch:
            fmt.Printf("%d^2 = %d\n", res.i, res.x)
        case <-timeout:
            fmt. Println("timeout.")
        }

    }
}

func timer(timer chan<- bool) {
    time.Sleep(1e9) // 1 second long
    timer <- true
}

func pow(i int, ch chan<- CalcResult) {
    x := i * i
    ch <- CalcResult{i, x}
}
