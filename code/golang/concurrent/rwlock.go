package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go write(2)
	go read(3)

	time.Sleep(5 * time.Second)
}

func read(i int) {
	println(i, "read start, reading lock")

	m.RLock()
	println(i, "[R] reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over, release lock")
}

func write(i int) {
	println(i, "write start, writing lock")

	m.Lock()
	println(i, "[W] writing")
	time.Sleep(1 * time.Second)
	m.Unlock()

	println(i, "write over, release lock")
}
