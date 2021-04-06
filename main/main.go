package main

import (
	"fmt"
	"goleetcode/solutions"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{}
	m := 3
	n := 0
	solutions.Merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
