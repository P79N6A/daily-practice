package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Luke, I'm your father.")
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
