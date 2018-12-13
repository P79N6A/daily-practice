package main

func main() {
	var a []int

	go func() { a = make([]int, 3) }()
	go func() { a = make([]int, 10) }()

	a[6] = 11
}
