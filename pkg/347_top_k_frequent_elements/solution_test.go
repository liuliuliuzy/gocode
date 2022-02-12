package topkfrequentelements

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{2, 2, 4, 4, 4, 4, 1, 1, 4, 1, 1, 6, 12, 56, 12, 3, 1, 3, 3, 31, 12, 1}
	ans := topKFrequent(nums, 3)
	correct := []int{1, 4, 3}
	for i := range correct {
		if ans[i] != correct[i] {
			t.Errorf("expected: %d, got: %d", correct[i], ans[i])
		}
	}
}
