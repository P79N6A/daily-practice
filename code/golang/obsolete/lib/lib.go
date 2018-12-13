package lib

import (
    "fmt"
)

var Holy string

func init() {
    Holy = "Holy Crap"
}

func Crap(str string) {
    fmt.Println(str)
}
