package tic_tac_toe

import "testing"

func TestFindWiner(t *testing.T) {

	tc := []struct {
		moves    [][]int
		expected string
	}{
		{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}, "A"},
		{[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}, "B"},
		{[][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}, "Draw"},
		{[][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}}, "B"},
		{[][]int{{0, 2}, {1, 0}, {2, 2}, {1, 2}, {2, 0}, {0, 0}, {0, 1}, {2, 1}, {1, 1}}, "A"},
	}

	for _, c := range tc {
		f := findWinner(c.moves)
		if f != c.expected {
			t.Errorf("expected %s, got %s", c.expected, f)
		}
	}
}

func findWinner(moves [][]int) string {
	rows := [3]int{0, 0, 0}
	cols := [3]int{0, 0, 0}
	diag, antiDiag := 0, 0

	for i, move := range moves {
		r, c := move[0], move[1]

		player := 1
		if i%2 == 1 {
			player = -1
		}

		rows[r] += player
		cols[c] += player

		if r == c {
			diag += player
		}
		if r+c == 2 {
			antiDiag += player
		}

		if abs(rows[r]) == 3 || abs(cols[c]) == 3 ||
			abs(diag) == 3 || abs(antiDiag) == 3 {
			if player == 1 {
				return "A"
			}
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
