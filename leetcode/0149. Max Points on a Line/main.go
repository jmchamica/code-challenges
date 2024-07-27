
func maxPoints(points [][]int) int {

	max := 1

	for i, _ := range points {
		x1 := points[i][0]
		y1 := points[i][1]

		for j := i + 1; j < len(points); j++ {
			x2 := points[j][0]
			y2 := points[j][1]

			count := 2
			a := y2 - y1
			b := x1 - x2
			c := a*x1 + b*y1
			// ax + by = c

			for k := j + 1; k < len(points); k++ {
				x3 := points[k][0]
				y3 := points[k][1]

				if a*x3+b*y3 == c {
					count++
				}
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}