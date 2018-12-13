package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) other(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "not found", http.StatusNotFound)
}

func (db database) header(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	for k, v := range req.Header {
		buf.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	fmt.Fprintf(w, "%s\n", buf.String())
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/header", db.header)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/", db.other)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
