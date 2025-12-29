package ice_cream_parlor

import (
	"fmt"
	"testing"
)

func TestIceCreamParlor(t *testing.T) {
	menu := []int{2, 7, 13, 5, 4, 13, 3}
	money := 10

	indices := getIndices(menu, money)
	fmt.Printf("%+v", indices)
}

func getIndices(menu []int, money int) []int {
	seen := make(map[int]int)
	for i, price := range menu {
		complement := money - price
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[price] = i
	}
	return []int{}
}
