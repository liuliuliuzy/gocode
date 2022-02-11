package minimumdifferencebetweenhighestandlowestofkscores

import (
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	// fmt.Println("after sort, ", nums)
	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i <= len(nums)-k; i++ {
		tmp := nums[i+k-1] - nums[i]
		if tmp < ans {
			ans = tmp
		}
	}
	return ans
}

func MinimumDifference(nums []int, k int) int {
	return minimumDifference(nums, k)
}
