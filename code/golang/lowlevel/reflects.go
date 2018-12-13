package main

import (
	"log"
	"reflect"
)

type Crapper struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func bindData(ptr interface{}) {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()

	log.Println("⇛ ", typ)
	log.Println("⇛ ", val)

	for i, n := 0, typ.NumField(); i < n; i++ {
		f := typ.Field(i)
		log.Println("Field: ", f.Name, f.Type)
		v := val.Field(i)
		log.Printf("Value: %q\n", v)
	}
}

func main() {
	holy := Crapper{ID: 1, Name: "Diddy"}

	bindData(&holy)
}
