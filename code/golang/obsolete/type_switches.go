package main

import "fmt"

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NIL"
	case int, uint:
		return fmt.Sprintf("%T: %d", x, x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	var v uint = 123
	fmt.Println(sqlQuote(v))
}
