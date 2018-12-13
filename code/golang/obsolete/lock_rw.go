package main

import (
	"log"
	"sync"
	"time"
)

var (
	mu      sync.RWMutex
	balance int
	wg      sync.WaitGroup
)

func Deposite(amount int) {
	wg.Add(1)
	mu.Lock()
	defer mu.Unlock()
	log.Println("W LOCK")

	balance += amount
	wg.Done()
}

// R 会阻塞 W
func Balance() {
	wg.Add(1)
	mu.RLock()
	defer mu.RUnlock()
	log.Println("R LOCK")

	time.Sleep(2 * time.Second)
	log.Println("BALANCE:", balance)
	wg.Done()
}

func main() {
	log.Println("INITIAL:", balance)

	go Balance()
	go Balance()
	go Deposite(100)

	time.Sleep(100 * time.Millisecond)
	wg.Wait()
	log.Println("END", balance)
}
