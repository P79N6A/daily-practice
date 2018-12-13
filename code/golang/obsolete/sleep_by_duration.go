package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Bars struct {
	Value string
	Index int
}

func (b *Bars) Next() byte {
	if (b.Index + 1) < len(b.Value) {
		b.Index++
	} else {
		b.Index = 0
	}
	return b.Value[b.Index]
}

func cleanup() {
	log.Println("AWAKEN.")
	os.Exit(0)
}

func main() {
	duration := flag.Duration("d", 1*time.Second, "tick duration")
	flag.Parse()

	// receive SIGNALs from OS
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
	}()

	log.Println("SLEEPING...")

	bars := Bars{`-\|/`, 0}
	blocker := time.Tick(*duration)
	// time.Tick made a wrapper with ticking
	for _ = range blocker {
		fmt.Printf("\r%c", bars.Next())
	}
}
