package main

import (
	"fmt"

	"github.com/rockdragon/daily-practice/code/golang/interview/DFA"
)

func main() {
	dfa := DFA.InitDFA([]*DFA.Status{
		DFA.InitStatus("0", "A", "B", "4"),
		DFA.InitStatus("1", "A", "B", "3"),
		DFA.InitStatus("2", "A", "B", "2"),
		DFA.InitStatus("3", "A", "B", "1"),
		DFA.InitStatus("4", "A", "B", "5"),
		DFA.InitStatus("5", "A", "B", "0"),
		DFA.InitStatus("6", "A", "B", "2"),
	})
	dfa.Run([]string{"A"})
	fmt.Println(dfa)
}
