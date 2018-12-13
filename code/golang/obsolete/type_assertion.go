package main

import "fmt"

type Greeting interface {
	Hello() string
}
type Calculator interface {
	Add(x, y int) int
}
type Nothing interface{}

// implement interface:
type HolyHigh struct{}

func (h *HolyHigh) Hello() string {
	return "hello"
}
func (h *HolyHigh) Add(x, y int) int {
	return x + y
}

func main() {
	var h Nothing
	h = new(HolyHigh)

	g := h.(Greeting)
	fmt.Println(g.Hello())

	c := h.(Calculator)
	fmt.Println(c.Add(10, 22))
}
