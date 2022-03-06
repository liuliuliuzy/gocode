package maximumdifferencebetweenincreasingelements

import "math"

func maximumDifference(nums []int) int {
	// l, r := math.MaxInt, 0
	// for _, n := range nums {
	// 	if n > l && n > r {
	// 		r = n
	// 	} else {
	// 		if l > n {
	// 			l = n
	// 		}
	// 	}
	// }
	// if l >= r {
	// 	return -1
	// }
	// return r - l

	// 维护左右两个数字是不对的，你无法确保一定是小的数字在左而大的数字在右
	// 正确做法应该是维护当前最大的差值与最小数字

	l, ans := math.MaxInt, -1
	for _, n := range nums {
		if n > l { // 一定要大于才去计算，这样才满足题目的条件
			ans = max(ans, n-l)
		} else {
			l = n
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumDifference(nums []int) int {
	return maximumDifference(nums)
}
