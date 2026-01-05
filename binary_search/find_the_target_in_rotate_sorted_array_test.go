package binary_search

import "testing"

func TestFindTheTargetInARotateSortedArray(t *testing.T) {
	in := []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	target := 1
	out := 2
	result := findTheTargetInARotateSortedArray(in, target)
	if result != out {
		t.Error("Expected ", out, " but got ", result)
	}
}

func findTheTargetInARotateSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
