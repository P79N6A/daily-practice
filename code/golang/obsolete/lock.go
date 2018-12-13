package main

import (
	"log"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposite(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func DisplayAccount() {
	log.Println("Account Balance: ", Balance())
}

func main() {
	go Deposite(10)
	go Deposite(20)
	go Deposite(30)
	go DisplayAccount()

	time.Sleep(2 * time.Second)
}
