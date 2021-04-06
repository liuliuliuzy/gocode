package main

import (
	"fmt"
	"goleetcode/solutions"
)

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3, 4}
	res := solutions.RemoveDuplicates(nums1)
	fmt.Println(res)
	for i := 0; i < res; i++ {
		fmt.Println(nums1[i])
	}
}
