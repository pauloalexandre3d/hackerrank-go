package maps_and_sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSumUnsorted(t *testing.T) {
	tcs := []struct {
		in     []int
		target int
		out    []int
	}{
		{[]int{-1, 3, 4, 2}, 3, []int{0, 2}},
	}
	for _, tc := range tcs {
		got := pairSumUnsorted(tc.in, tc.target)
		if !assert.Equal(t, tc.out, got) {
			t.Errorf("input: %v, got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func pairSumUnsorted(in []int, target int) []int {
	m := make(map[int]int)
	for i, x := range in {
		if _, exists := m[target-x]; exists {
			return []int{m[target-x], i}
		}
		m[x] = i
	}
	return []int{}
}
