package main

import (
	"log"
	"time"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposite(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
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
