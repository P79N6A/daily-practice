package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{
"name": {"first":"Janet","last":"Prichard"},
"age":47 
}`

func main() {
	firstName := gjson.Get(json, "name.first")
	fmt.Println(firstName)
}
