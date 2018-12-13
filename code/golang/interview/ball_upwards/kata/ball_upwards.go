// https://www.codewars.com/kata/ball-upwards/go
package kata

func MaxBall(v0 int) (times int) {
	var max float64
	v := float64(v0) * 1000 / 3600
	g := 9.81

	// h(s) = v0 * t + g * t^2 / 2
	for t := 0.0; ; t += 0.1 {
		h := v*t - 0.5*g*t*t
		if h > max {
			max = h
			times++
		} else if h < max {
			break
		}
	}

	return
}
