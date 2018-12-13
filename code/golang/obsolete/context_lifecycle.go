package main

import (
	"context"
	"log"
	"time"
)

func taskA(ctx context.Context) {
	ctxA, cancel := context.WithTimeout(ctx, 2*time.Second)
	ch := make(chan int)
	defer cancel()

	go taskB(ctxA, ch)

	for {
		select {
		case <-ctx.Done():
			log.Println("testA Done")
			return
		case <-ch:
		}
	}
}

func taskB(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			log.Println("testB Done")
			close(ch)
			return
		default:
			ch <- 1
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go taskA(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
