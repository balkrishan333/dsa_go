package main

import (
	"fmt"
	"math"
)

func minCostConnectPoints(points [][]int) int {
	length := len(points)
	visited := make([]bool, length)

	curr, result, edges := 0, 0, 0
	visited[0] = true
	distArr := make([]int, length)
	for i := 0; i < length; i++ {
		distArr[i] = math.MaxInt
	}

	for edges < length-1 {
		next := -1
		minEdge := math.MaxInt

		for i := 0; i < length; i++ {
			if !visited[i] {
				cost := Abs(points[curr][0]-points[i][0]) + Abs(points[curr][1]-points[i][1])
				distArr[i] = Min(distArr[i], cost)

				if distArr[i] < minEdge {
					minEdge = distArr[i]
					next = i
				}
			}
		}
		result += minEdge
		curr = next
		edges++
		visited[curr] = true
	}
	return result
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	points := [][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}
	fmt.Println(minCostConnectPoints(points))
}
