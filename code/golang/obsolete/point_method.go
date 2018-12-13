package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Mail string
}

type Team map[string]Person

func (t *Team) Add(p Person) {
	(*t)[p.Name] = p
}
func (t *Team) Get(name string) Person {
	return (*t)[name]
}

func main() {
	t := make(Team)
	p1 := Person{"Tom", 10, "tom@tom.com"}
	t.Add(p1)
	fmt.Println(t.Get("Tom"))
}
