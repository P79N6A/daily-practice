package main

import "fmt"
import "reflect"
import "math/rand"
import "time"
import "os"

func main() {
    var i, j, k int
    fmt.Println(i, j, k)

    w := "hello"
    fmt.Printf("%s type: %s\n", w, reflect.TypeOf(w))

    rand.Seed(time.Now().UnixNano())
    fmt.Println("random number:", rand.Float64())

    f, err := os.Open("/etc/hosts")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)
    f.Close()

    x := 1
    p1 := &x
    p2 := new(int)
    fmt.Println("pointer:", p1, *p2)
}
