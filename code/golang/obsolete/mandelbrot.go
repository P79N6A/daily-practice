package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "flag"
    "fmt"
)

var filename = flag.String("f", "./mandelbrot.png", "png file name.")

func main() {
    // construction
    flag.Parse()
    fd, err := os.Create(*filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    // computation
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 1024, 1024
    )
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(fd, img) // NOTE: ignoring errors
    fmt.Println("DONE")
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15
    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        coefficient := contrast*n
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - coefficient}
            // return color.RGBA{ 25, 104, 17, 255 - coefficient}
        }
    }
    return color.Black
}
