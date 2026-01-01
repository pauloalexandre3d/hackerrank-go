package maps_and_sets

import (
	"fmt"
	"testing"
)

func TestZeroStriping(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 0, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 0},
	}

	m, n := len(matrix), len(matrix[0])

	zeroRows := make(map[int]struct{})
	zeroCols := make(map[int]struct{})

	for r := range m {
		for c := range n {
			if matrix[r][c] == 0 {
				zeroRows[r] = struct{}{}
				zeroCols[c] = struct{}{}
			}
		}
	}
	for r := range m {
		for c := range n {
			_, rowHasZero := zeroRows[r]
			_, colHasZero := zeroCols[c]
			if rowHasZero || colHasZero {
				matrix[r][c] = 0
			}
		}
	}
	for _, row := range matrix {
		fmt.Println(row)
	}
}
