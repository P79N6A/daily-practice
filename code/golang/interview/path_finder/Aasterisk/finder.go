// http://blog.jobbole.com/71044/
// https://www.redblobgames.com/pathfinding/a-star/implementation.html
package Aasterisk

type Finder struct {
	start    Coordinate
	target   Coordinate
	barriers []Coordinate
	width    int
	height   int
}

func (f *Finder) InBounds(n Coordinate) bool {
	return n.X >= 0 && n.X < f.width && n.Y >= 0 && n.Y < f.height
}

func (f *Finder) Passable(n Coordinate) bool {
	for _, v := range f.barriers {
		if v.sameAs(n) {
			return false
		}
	}
	return true
}

func (f *Finder) HasArrived(n Coordinate) bool {
	return f.target.sameAs(n)
}

func (f *Finder) Neighbors(p Coordinate) (neighbors []Coordinate) { // vertically & horizontally
	coordinates := []Coordinate{
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
	}

	for _, pt := range coordinates {
		if f.InBounds(pt) && f.Passable(pt) {
			neighbors = append(neighbors, pt)
		}
	}

	return
}

func (f *Finder) MoveCosts(from, to Coordinate) int {
	return 1
}

func (f *Finder) Search() (cameFrom map[Coordinate]interface{}, costSoFar map[Coordinate]int) {
	compare := func(i, j int) bool { return i < j }
	frontier := NewPriorityQueue(NewItem(f.start, 0), compare)

	cameFrom = make(map[Coordinate]interface{})
	cameFrom[f.start] = nil
	costSoFar = make(map[Coordinate]int)
	costSoFar[f.start] = 0

	for frontier.Len() > 0 {
		item := frontier.Remove()
		current := item.Value.(Coordinate)

		if f.HasArrived(current) {
			break
		}

		for _, next := range f.Neighbors(current) {
			newCost := costSoFar[current] + f.MoveCosts(current, next)
			vNext, ok := costSoFar[next]
			if !ok || newCost < vNext {
				costSoFar[next] = newCost
				priority := newCost + next.Manhattan(f.target)
				frontier.Insert(NewItem(next, priority))
				cameFrom[next] = current
			}
		}
	}

	return
}

func NewFinder(start, target Coordinate, barriers []Coordinate, w, h int) *Finder {
	return &Finder{
		start:    start,
		target:   target,
		barriers: barriers,
		width:    w,
		height:   h,
	}
}
