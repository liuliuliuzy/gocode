package base7

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	for i := -20; i < 22; i++ {
		fmt.Println(i, convertToBase(i))
	}
	fmt.Println(convertToBase(1e7))
}
