package matrices

import (
	"fmt"
	"testing"
)

func TestMatrices(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			v := grid[r][c]
			fmt.Printf("r: %d, c: %d, v: %d\n", r, c, v)
		}
	}
}
