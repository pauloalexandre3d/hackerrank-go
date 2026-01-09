package best_fit_decreasing

import (
	"fmt"
	"slices"
	"testing"
)

type node struct {
	sumCpusRequests int
	cpusRequests    []int
}

func (n node) String() string {
	return fmt.Sprintf("Total : %d, cpus requests: %+v", n.sumCpusRequests, n.cpusRequests)
}

func TestBestFirDecreasing(t *testing.T) {
	//in := []int{6, 7, 10, 1, 4}
	in := []int{3, 2, 5, 4, 1, 6, 2, 3}
	maxCpusRequests := 10
	//out := [][]int{
	//	{6, 4},
	//	{7, 1},
	//	{10},
	//}

	slices.Sort(in)
	slices.Reverse(in)
	fmt.Println(in)

	nodes := make([]*node, 0)

	for _, cpuReq := range in {

		bestIdx := -1
		bestSum := -1

		for i, node := range nodes {
			if node.sumCpusRequests+cpuReq <= maxCpusRequests &&
				node.sumCpusRequests > bestSum {

				bestIdx = i
				bestSum = node.sumCpusRequests
			}
		}

		if bestIdx != -1 {
			nodes[bestIdx].cpusRequests = append(nodes[bestIdx].cpusRequests, cpuReq)
			nodes[bestIdx].sumCpusRequests += cpuReq
		} else {
			nodes = append(nodes, &node{
				cpusRequests:    []int{cpuReq},
				sumCpusRequests: cpuReq,
			})
		}
	}

	fmt.Printf("nodes: %+v\n", nodes)

}
