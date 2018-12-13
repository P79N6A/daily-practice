package main

import (
	"log"
	"reflect"
	"strings"
)

type requestStruct struct {
	Labels     []string            `http:"l"`
	MaxResults int                 `http:"max"`
	Exact      bool                `http:"x"`
	Holy       func(string) string `http:"POST"`
}

func extract(ptr interface{}) error {
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}
	log.Println(fields)

	t := reflect.ValueOf(ptr).Type()
	for i := 0; i < t.NumMethod(); i++ {
		methodInfo := t.Method(i)
		log.Println("func " + methodInfo.Name + "()")
	}

	return nil
}

func main() {
	var data requestStruct
	extract(&data)

}
