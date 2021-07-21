package solutions

import (
	"sort"
)

//证了一下，排序之后，必然是最小的和最大的、第二小的和第二大的、...这样搭配能够使得结果最小
func minPairSum(nums []int) int {
	res := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		res = max(res, nums[i]+nums[n-1-i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinPairSum(nums []int) int {
	return minPairSum(nums)
}
