package minimumcosttoreachdestinationintime

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	maxTime := 29
	edges := [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}
	passingFees := []int{5, 1, 2, 20, 20, 3}
	fmt.Println(minCost(maxTime, edges, passingFees))
}
