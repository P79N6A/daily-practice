package Aasterisk

type Coordinate struct {
	X, Y int
}

func (n Coordinate) distanceX(a Coordinate) int {
	return distance(n.X, a.X)
}
func (n Coordinate) distanceY(a Coordinate) int {
	return distance(n.Y, a.Y)
}
func (n Coordinate) sameAs(a Coordinate) bool {
	return n.X == a.X && n.Y == a.Y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
