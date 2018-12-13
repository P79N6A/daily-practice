// https://www.codewars.com/kata/559ce00b70041bc7b600013d/train/go
package kata

import "fmt"

func Finance(n int) int {
	return n * (n + 1) * (n + 2) / 2
}

func main() {
	fmt.Println(Finance(5))
	fmt.Println(Finance(6))
	fmt.Println(Finance(7))
	fmt.Println(Finance(8))
	fmt.Println(Finance(15))
}
