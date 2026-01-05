package binary_search

import "testing"

func TestFindTheInsertionIndex(t *testing.T) {
	tc := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 2, 4, 5, 7, 8, 9},
			target: 4,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   3,
		},
	}

	for _, tt := range tc {
		got := findTheInsertionIndex(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("searchInsert() = %v, want %v", got, tt.want)
		}
	}
}

func findTheInsertionIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
