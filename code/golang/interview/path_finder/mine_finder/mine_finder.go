// K个矿点，使用 bfs 搜索查找轮数最小的基地点
package mine_finder

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Equals(pt Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type MineField struct {
	width, height int
	mines         []Point
	visited       map[Point]bool
}

func (mf *MineField) validPoint(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < mf.width && p.Y < mf.height
}
func (mf *MineField) visit(p Point) {
	mf.visited[p] = true
}
func (mf *MineField) hasVisited(p Point) bool {
	_, ok := mf.visited[p]
	return ok
}
func (mf *MineField) clearVisited() {
	mf.visited = make(map[Point]bool)
}
func (mf *MineField) IsMine(p Point) bool {
	for _, mine := range mf.mines {
		if mine.Equals(p) {
			return true
		}
	}
	return false
}

func (mf *MineField) Neighbors(p Point) (neighbors []Point) { // vertically & horizontally
	pts := []Point{
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
	}

	for _, pt := range pts {
		if mf.validPoint(pt) {
			neighbors = append(neighbors, pt)
		}
	}

	return
}

func (mf *MineField) search(queue []Point, found int, round int) int { // BFS
	if len(queue) == 0 || found == len(mf.mines) {
		return round
	}
	round++
	nextRound := make([]Point, 0)
	for _, pt := range queue {
		if !mf.hasVisited(pt) {
			mf.visit(pt)
			if mf.IsMine(pt) {
				found++
			}
			for _, neigh := range mf.Neighbors(pt) {
				if !mf.hasVisited(neigh) {
					nextRound = append(nextRound, neigh)
				}
			}
		}
	}
	return mf.search(nextRound, found, round)
}

func (mf *MineField) Search() (minPoint []Point, minRound int) {
	minRound = math.MaxInt32

	for x := 0; x < mf.width; x++ {
		for y := 0; y < mf.height; y++ {
			current := Point{x, y}
			if mf.IsMine(current) {
				continue
			}
			round := mf.search([]Point{current}, 0, 0)
			if round < minRound {
				minRound = round
				minPoint = []Point{current}
			} else if round == minRound {
				minPoint = append(minPoint, current)
			}
			mf.clearVisited()
		}
	}

	return
}

func NewMineField(w, h int, mines []Point) *MineField {
	return &MineField{
		width:   w,
		height:  h,
		mines:   mines,
		visited: make(map[Point]bool),
	}
}
