package main

import (
    "fmt"
    "image/color"
)

type Point struct { X, Y float64 }
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

type ColoredPoint struct {
    Point
    Color color.RGBA
}

func main() {
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 255, 0, 255}
    p := ColoredPoint{Point{1, 2}, red}
    q := ColoredPoint{Point{4, 5}, blue}
    p.ScaleBy(2)
    q.ScaleBy(2)
    fmt.Println(p, q)
}
