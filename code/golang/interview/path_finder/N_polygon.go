// https://blog.csdn.net/jinguangliu/article/details/78745012
package path_finder

import "fmt"

type Coordinate struct {
	X, Y int
}

func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
func (c Coordinate) absX(a Coordinate) int {
	return abs(c.X, a.X)
}
func (c Coordinate) absY(a Coordinate) int {
	return abs(c.Y, a.Y)
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Polygon struct {
	vertexes []Coordinate
}

func (p *Polygon) Perimeter() (perimeter int) {
	for i, v := range p.vertexes {
		next := p.vertexes[i+1]
		if v.X == next.X {
			perimeter += v.absY(next)
		} else {
			perimeter += v.absX(next)
		}
	}
	return
}

func (p *Polygon) Divides(k uint) (divides []Coordinate) {
	perimeter := p.Perimeter()
	divLen := perimeter / int(k)
	length := divLen
	index := 0

	for i, current := range p.vertexes {
		np := i + 1
		if i == len(p.vertexes)-1 {
			np = 0
		}
		next := p.vertexes[np]

		if current.X == next.X {
			distance := current.absY(next)
			if length <= distance {
				base := 0
				for length <= distance {
					if current.Y > next.Y {
						divides[index] = Coordinate{current.X, current.Y - (length + base)}
					} else {
						divides[index] = Coordinate{current.X, current.Y + (length + base)}
					}

					base += length
					distance -= length
					length = divLen
					index++
				}

				length = divLen - distance
			} else {
				length -= distance
			}
		} else {
			distance := current.absY(next)
			if length <= distance {
				base := 0
				for length <= distance {
					if current.X > next.X {
						divides[index] = Coordinate{current.X - (length + base), current.Y}
					} else {
						divides[index] = Coordinate{current.X + (length + base), current.Y}
					}

					base += length
					distance -= length
					length = divLen
					index++
				}

				length = divLen - distance
			} else {
				length -= distance
			}
		}
	}
	return
}
