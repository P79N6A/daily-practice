package main

import (
	"log"
	"reflect"
	"unsafe"
)

type Job struct {
	Company string
	Title   string
}
type Person struct {
	Name string
	Age  uint8
	Bio  string
	Jobs []Job
}

func printLayout(ptr interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		log.Printf("[%s]:\t Offset: %d\n", fieldInfo.Name, fieldInfo.Offset)
	}
	return
}

func main() {
	// 1. enumerate its fields
	var man Person
	printLayout(&man)

	// 2. ptr operation
	pa := (*uint8)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&man)) + unsafe.Offsetof(man.Age)))
	*pa = 66
	log.Println("man.Age = ", man.Age)
}
