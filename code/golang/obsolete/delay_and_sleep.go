package main

import (
    "fmt"
    "time"
)

type Rocket struct { Name string }
func (r *Rocket) Launch() {
    fmt.Printf("%s %s was launched.\n", now(), r.Name)
}

func now() string {
    return fmt.Sprintf("[%s]", time.Now().Format(time.RFC3339))
}

func main() {
    r := Rocket{"Apollo I"}
    fmt.Printf("%s the astronaut was asleep.\n", now())
    time.AfterFunc(10 * time.Second, r.Launch)
    fmt.Printf("%s at meanwhile, the rocket was ready to launch.\n", now())
    time.Sleep(20 * time.Second)

    defer func() {
        fmt.Printf("%s the astronaut was waken up.\n", now())
    }()
}
