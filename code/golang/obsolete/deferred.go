package main

import (
    "fmt"
)

func rangeOf(begin, end int, f func(curr int)) []int{
    var result []int
    for i := begin; i< end; i++ {
        curr := i
        f(curr)
        result = append(result, curr)
    }
    return result
}

func task() []int {
    fmt.Println("holy start.")
    indexes := rangeOf(1, 100, func(curr int) {
        fmt.Println("holy process:", curr)
    })
    
    defer func() {
        fmt.Println("holy end.")
    }()
    return indexes
}

func main() {
    indexes := task()
    fmt.Println("indexes:", indexes)
}
