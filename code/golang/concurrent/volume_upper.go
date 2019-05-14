package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var NUM int = 1048575

func schedule() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("TOTAL: ", runtime.NumGoroutine())
		}
	}()

	for i := 0; i < NUM; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		runtime.Gosched()
	}
}

func main() {
	wg.Add(NUM)

	go schedule()

	wg.Wait()
	fmt.Println("DONE")
}
