package main

import (
	"log"

	frontend "./frontend"
)

var (
	expression = "A(B(,C(,)),D(E(,),F(,)))"
)

func main() {
	log.Println("INPUT:", expression)
	tree := frontend.NewParser(expression)
	node := tree.ParseNode()
	log.Println("RESULT: ", node)

	log.Println(frontend.ID)
}
