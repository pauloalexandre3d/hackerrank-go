package quicksort

import (
	"fmt"
	"testing"
)

//func main() {
////Enter your code here. Read input from STDIN. Print output to STDOUT
//scanner := bufio.NewScanner(os.Stdin)
//// Read first line: array size
//scanner.Scan()
//n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
//
//// Read second line: array numbers
//scanner.Scan()
//l := strings.Fields(scanner.Text())
//arr := make([]int, n)
//for i, s := range l {
//	arr[i], _ = strconv.Atoi(s)
//}

//quicksort(arr, 0, len(arr) - 1)
//}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 9, 8, 2, 7, 5}
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, left, right int) {
	// fmt.Printf("arr: %v left: %d right: %d\n", arr, left, right)
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right)

	printArray(arr)

	quicksort(arr, left, pivotIndex-1)
	quicksort(arr, pivotIndex+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right] // LAST element as pivot
	i := left - 1       // boundary for elements < pivot

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			// Swap elements < pivot to left side
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place pivot in correct position
	i++
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Print(v)
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
