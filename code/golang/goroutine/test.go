package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errChan := make(chan string)

	go func(ctx context.Context, errChan chan<- string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				errChan <- "holy crap"
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx, errChan)

	time.Sleep(2 * time.Second)

	cancel()
	fmt.Println(ctx.Err())

	for {
		select {
		case err := <-errChan:
			fmt.Println("监控退出原因：", err)
			return
		}
	}
}
