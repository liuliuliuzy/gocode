package findtheminimumnumberoffibonaccinumberswhosesumisk

import "testing"

func TestSolution(t *testing.T) {
	got := findMinFibonacciNumbers(7)
	if got != 2 {
		t.Error("wrong")
		t.Errorf("findMinFibonacciNumbers(7) = %d", got)
		t.Error("excepted: 2")
	}
}
