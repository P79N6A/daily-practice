package main

import (
	. ".."
	"fmt"
)

func main() {
	w, h := 10, 10
	mine := []Point{
		{5, 3},
		{7, 1},
		{7, 6},
	}
	//mine := []Point{
	//	{9, 6},
	//	{6, 9},
	//	{9, 9},
	//}

	// BFS search miner base
	mf := NewMineField(w, h, mine)

	fmt.Println(mf.Search())
}
