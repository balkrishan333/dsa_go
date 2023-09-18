package main

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {

	type RowStrength struct {
		strength, index int
	}
	rowStrengths := make([]RowStrength, len(mat))

	for i, row := range mat {
		strength := 0
		for _, val := range row {
			strength += val
		}
		rowStrengths[i] = RowStrength{strength, i}
	}

	sort.Slice(rowStrengths, func(i, j int) bool {
		if rowStrengths[i].strength == rowStrengths[j].strength {
			return rowStrengths[i].index < rowStrengths[j].index
		}
		return rowStrengths[i].strength < rowStrengths[j].strength
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = rowStrengths[i].index
	}
	return result
}

// func main() {
// 	k := 3
// 	mat := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 0},
// 		{1, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 1},
// 	}
// 	fmt.Println(kWeakestRows(mat, k))
// }
