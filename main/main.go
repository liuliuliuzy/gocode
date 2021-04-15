package main

import (
	"fmt"
	"goleetcode/solutions"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	res := solutions.FindMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
