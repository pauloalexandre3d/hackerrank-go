package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSum(t *testing.T) {
	input := []int{-5, -2, 3, 4, 6}
	target := 7

	result := pairSumSorted(input, target)

	assert.Equal(t, []int{2, 3}, result)
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
