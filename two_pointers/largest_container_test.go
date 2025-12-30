package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestContainer(t *testing.T) {
	tcs := []struct {
		heights []int
		output  int
	}{
		{[]int{}, 0},
		{[]int{2, 7, 8, 3, 7, 6}, 24},
	}
	for _, c := range tcs {
		got := largestContainer(c.heights)
		if !assert.Equal(t, c.output, got) {
			t.Errorf("input: %v, got %v, want %v", c.heights, got, c.output)
		}
	}
}

func largestContainer(heights []int) interface{} {
	maxWater := 0
	left, right := 0, len(heights)-1
	for left < right {
		water := min(heights[left], heights[right]) * (right - left)
		maxWater = max(maxWater, water)

		if heights[left] < heights[right] {
			left++
		} else if heights[left] > heights[right] {
			right--
		} else {
			left++
			right--
		}
	}
	return maxWater
}
