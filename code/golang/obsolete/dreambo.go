package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Printf("URL.path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}
