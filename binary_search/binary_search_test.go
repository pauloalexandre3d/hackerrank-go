package binary_search

import "testing"

func TestName(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 13, 15, 17, 19, 23, 25}
	x := 9

	r := false

	left := 0
	right := len(a)
	for {
		if left <= right {
			mid := left + ((right - left) / 2)
			if a[mid] == x {
				r = true
				break
			} else if x < a[mid] {
				right = mid - 1
				//r = binarySearch(a, x, left, mid-1)
			} else {
				left = mid + 1
				//r = binarySearch(a, x, mid+1, right)
			}
		}
	}

	println(r)
}

func binarySearch(a []int, x, left, right int) bool {
	if left > right {
		return false
	}
	mid := left + ((right - left) / 2)
	if a[mid] == x {
		return true
	} else if x < a[mid] {
		return binarySearch(a, x, left, mid-1)
	}

	return binarySearch(a, x, mid+1, right)
}
