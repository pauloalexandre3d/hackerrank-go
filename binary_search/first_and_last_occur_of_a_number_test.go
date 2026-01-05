package binary_search

import "testing"

func TestFirstAndLastOccurrencesOfANumber(t *testing.T) {
	in := []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 9, 10, 11}
	target := 4
	expected := []int{3, 5}

	lb := lowerBound(in, target)
	ub := upperBound(in, target)
	if lb != expected[0] || ub != expected[1] {
		t.Errorf("Expected %v, got %v", expected, []int{lb, ub})
	}
}

func upperBound(in []int, target int) int {
	left, right := 0, len(in)-1

	for left < right {
		mid := (left + right) / 2
		if in[mid] > target {
			right = mid - 1
		} else if in[mid] < target {
			left = mid + 1
		} else {
			left = mid + 1
		}
	}
	if in[right] == target {
		return right
	}
	return -1
}

func lowerBound(in []int, target int) int {
	left, right := 0, len(in)-1

	for left < right {
		mid := (left + right) / 2
		if in[mid] > target {
			right = mid - 1
		} else if in[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if in[left] == target {
		return left
	}
	return -1
}
