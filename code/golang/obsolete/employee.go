package main

import (
    "fmt"
    "time"
)

type Employee struct {
    ID          int
    Name        string
    Address     string
    DoB         time.Time
    Position    string
    Salary      int
    ManagerID   int
}

func main() {
    var dilbert Employee

    fmt.Println(dilbert.ID)
}
