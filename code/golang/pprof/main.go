package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/rockdragon/daily-practice/code/golang/pprof/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/rockdragon"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
