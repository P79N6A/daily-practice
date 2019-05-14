package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	num := 9876

	for i := 0; i < num; i++ {
		go func(idx int) {
			fmt.Println(idx)
			fmt.Println(runtime.NumGoroutine())
		}(i)
		runtime.Gosched()
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
