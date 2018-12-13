package main

import (
	. ".."
)

func main() {
	start := Coordinate{0, 0}
	target := Coordinate{0, 18}
	width, height := 19, 19
	//barriers := []Coordinate{
	//	{7, 8}, {7, 9}, {7, 10},
	//	{8, 11}, {9, 12}, {10, 12}, {11, 11},
	//	{12, 2}, {12, 3}, {12, 4}, {12, 5}, {12, 6},
	//	{12, 7}, {12, 8}, {12, 9}, {12, 10},
	//}
	barriers := []Coordinate{
		{7, 8}, {7, 9}, {7, 10},
		{8, 11}, {9, 12}, {10, 12}, {11, 11},
		{12, 0}, {12, 1}, {12, 2}, {12, 3}, {12, 4},
		{12, 5}, {12, 6}, {12, 7}, {12, 8}, {12, 9}, {12, 10},
	}
	finder := NewFinder(start, target, barriers, width, height)

	cameFrom, _ := finder.Search()
	if from, ok := cameFrom[target]; ok { //construct reverse path
		Log(target)
		for from != nil {
			Log(from)
			from = cameFrom[from.(Coordinate)]
		}
	}
}
