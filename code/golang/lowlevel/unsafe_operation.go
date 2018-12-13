package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	x float64
}

func (f Foo) BBB() float64 {
	return f.x
}

type FX interface {
	BBB() float64
}

type MyInterface struct {
	Type  *Foo
	Value *float64
}

func main() {
	var foo FX = Foo{x: 100.0}
	pfoo := (*MyInterface)(unsafe.Pointer(&foo))
	fmt.Printf("%f\n", *(*pfoo).Value)
}
