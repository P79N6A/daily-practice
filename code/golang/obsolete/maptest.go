package main

import "log"

// ConnMap ...
type ConnMap map[string]int

func NewConnMap() (c ConnMap) {
	return make(ConnMap)
}

func (c ConnMap) Delete(k string) {
	_, ok := c[k]
	if ok {
		delete(c, k)
	}
}

func (c ConnMap) Get(k string) (elem int, ok bool) {
	elem, ok = c[k]
	return
}

func main() {
	c := NewConnMap()
	c["a"] = 10
	c["b"] = 20
	c.Delete("b")
	log.Println(c)
}
