package main

import (
    "fmt"
    "math"
)

type Point struct { x, y float64 }

// method version
func (this Point) Distance(p Point) float64 {
    return math.Hypot(p.x-this.x, p.y-this.y)
}
// function version
func Distance(p2 Point, p1 Point) float64 {
    return math.Hypot(p2.x-p1.x, p2.y-p1.y)
}

func main() {
    p1 := Point{1, 4}
    p2 := Point{3, 10}
    fmt.Println(Distance(p2, p1))
    fmt.Println(p2.Distance(p1))
}
