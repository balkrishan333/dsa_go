package main

import "math"

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	print(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {

	ans := 0

	for i := 0; i < len(points)-1; i++ {
		currX := points[i][0]
		currY := points[i][1]

		targetX := points[i+1][0]
		targetY := points[i+1][1]

		ans += int(math.Max(math.Abs(float64(currX-targetX)), math.Abs(float64(currY-targetY))))
	}
	return ans
}
