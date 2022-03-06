package pathwithminimumeffort

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
}
