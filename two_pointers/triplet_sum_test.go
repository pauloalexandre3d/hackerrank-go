package two_pointers

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripletSum(t *testing.T) {
	tc := []struct {
		input []int
		want  [][]int
	}{
		{
			input: []int{-3, 0, 1, 2, -1, 1, -2},
			want:  [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
		{
			input: []int{-5, 2, -1, -2, 3},
			want:  [][]int{{-5, 2, 3}, {-2, -1, 3}},
		},
	}
	for _, c := range tc {
		got := tripletSum(c.input)
		if !assert.Equal(t, got, c.want) {
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}

func tripletSum(input []int) [][]int {
	var triplets [][]int // Fix: correct type
	slices.Sort(input)

	for i := 0; i < len(input)-2; i++ { // Fix: correct loop
		if input[i] > 0 {
			break
		}
		if i > 0 && input[i] == input[i-1] {
			continue
		}
		pairs := pairSumSortedAllPairs(input, i+1, -input[i])
		for _, pair := range pairs {
			triplet := []int{input[i], pair[0], pair[1]}
			triplets = append(triplets, triplet)
		}
	}
	return triplets
}

func pairSumSortedAllPairs(input []int, start, target int) [][]int {
	var pairs [][]int
	left, right := start, len(input)-1
	for left < right {
		sum := input[left] + input[right]
		if sum == target {
			pairs = append(pairs, []int{input[left], input[right]})
			left++
			for left < right && input[left] == input[left-1] {
				left++
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return pairs
}
