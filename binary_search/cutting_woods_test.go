package binary_search

import (
	"slices"
	"testing"
)

func TestCuttingWoods(t *testing.T) {
	in := []int{2, 6, 3, 8}
	k := 7
	out := 3
	res := cuttingWoods(in, k)
	if res != out {
		t.Errorf("Expected %d, but got %d", out, res)
	}
}

func cuttingWoods(heights []int, k int) int {
	left, right := 0, slices.Max(heights)

	for left < right {
		mid := ((left + right) / 2) + 1
		if cutsEnoughWood(mid, k, heights) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func cutsEnoughWood(h, k int, heights []int) bool {
	woodCollected := 0
	for _, height := range heights {
		if height > h {
			woodCollected += height - h
		}
	}
	return woodCollected >= k
}
