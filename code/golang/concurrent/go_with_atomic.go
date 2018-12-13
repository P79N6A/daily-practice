package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var n int64
	var f func()
	f = func() {
		go f()
		atomic.AddInt64(&n, 1)
	}
	f()

	for range time.NewTicker(time.Second).C {
		fmt.Printf("%d\n", atomic.LoadInt64(&n))
	}
}
