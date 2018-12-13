package Aasterisk

import "math"

// 曼哈顿距离(L1)
// 在正方形网格中，允许向4邻域的移动
// 在六边形网格中，允许6个方向的移动，使用适合于六边形网格的曼哈顿距离
func (n Coordinate) Manhattan(a Coordinate) int {
	return n.distanceX(a) + n.distanceY(a)
}

// 切比雪夫距离(L∞)
// 在正方形网格中，允许向8邻域的移动
func (n Coordinate) Chebyshev(a Coordinate) int {
	return max(n.distanceX(a), n.distanceY(a))
}

// 欧几里得距离(L2)
// 在正方形网格中，允许任何方向的移动，可能适合，但也可能不适合
func (n Coordinate) Euclid(a Coordinate) float64 {
	d := math.Pow(float64(n.distanceX(a)), 2)
	d += math.Pow(float64(n.distanceY(a)), 2)
	return math.Sqrt(d)
}
