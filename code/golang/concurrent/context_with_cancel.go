package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan interface{})
	go produce(ctx, ch)
	go consume(ctx, ch)

	time.Sleep(5 * time.Second)
	cancel()
}

func produce(ctx context.Context, out chan<- interface{}) error {
	i := 0
	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		case out <-i:
			time.Sleep(100 * time.Millisecond)
		}
		i++
	}
}

func consume(ctx context.Context, in <-chan interface{}) error {
	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		case i := <- in:
			fmt.Println("i:", i)
		}
	}
}