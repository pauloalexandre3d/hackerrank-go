package maps_and_sets

import (
	"fmt"
	"testing"
)

func TestVerifySudokuBoard(t *testing.T) {
	invalidBoard := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},

		// duplicate 5 in this row
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 5, 8, 0, 3, 0, 0, 1}, // another 5 here (row conflict)

		{7, 0, 0, 0, 2, 0, 0, 0, 6},

		// duplicate 9 in this column (col 0)
		{9, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{9, 0, 0, 0, 8, 0, 0, 7, 9}, // 9 repeated in col 0 and in this row/box
	}

	result := verifySudokuBoard(invalidBoard)
	fmt.Println(result)
	if result {
		t.Error("Expected invalid sudoku, got valid")
	}
}

func verifySudokuBoard(board [][]int) bool {

	rowSets := make([]map[int]struct{}, 9)
	for i := range rowSets {
		rowSets[i] = make(map[int]struct{})
	}
	columnSets := make([]map[int]struct{}, 9)
	for i := range columnSets {
		columnSets[i] = make(map[int]struct{})
	}
	subgridSets := [3][3]map[int]struct{}{}

	// Initialize all 9 subgrids
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subgridSets[i][j] = make(map[int]struct{})
		}
	}

	for row := range 9 {
		for col := range 9 {
			num := board[row][col]
			if num == 0 {
				continue
			}

			if _, exists := rowSets[row][num]; exists {
				return false
			}
			rowSets[row][num] = struct{}{}

			if _, exists := columnSets[col][num]; exists {
				return false
			}
			columnSets[col][num] = struct{}{}

			if _, exists := subgridSets[row/3][col/3][num]; exists {
				return false
			}
			subgridSets[row/3][col/3][num] = struct{}{}
		}
	}
	return true
}
