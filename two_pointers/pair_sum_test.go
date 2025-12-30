package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	tc := []struct {
		in     []int
		target int
		want   []int
	}{
		{[]int{-5, -2, 3, 4, 6}, 7, []int{2, 3}},
		{[]int{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{[]int{2, 5, 9, 11}, 11, []int{0, 2}},
	}

	for _, c := range tc {
		got := pairSumSorted(c.in, c.target)
		assert.Equal(t, c.want, got)
	}
}

func pairSumSorted(input []int, target int) interface{} {
	left, right := 0, len(input)-1
	for left < right {
		sum := input[left] + input[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{}
}
