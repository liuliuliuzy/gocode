package solutions

import "math"

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
*/

// 贪心算法

// 不对，想得太简单了
// func increasingTriplet(nums []int) bool {
// 	// 找到左右两个端点
// 	l, r := 0, len(nums)-1
// 	for l < r && nums[l] > nums[l+1] {
// 		l++
// 	}
// 	for l < r && nums[r] < nums[r-1] {
// 		r--
// 	}
// 	// 只有这样才可能存在符合条件的三元组
// 	if l < r-1 && nums[l] < nums[r]-1 {
// 		for k := l + 1; k < r; k++ {
// 			if nums[k] < nums[r] && nums[k] > nums[l] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// 解法一
// 辅助数组 双向遍历
// func increasingTriplet(nums []int) bool {
// 	n := len(nums)
// 	if n < 3 {
// 		return false
// 	}
// 	leftMin := make([]int, n)
// 	rightMax := make([]int, n)
// 	fmt.Println(leftMin)
// 	fmt.Println(rightMax)
// 	for i := 0; i < n; i++ {
// 		if i == 0 {
// 			leftMin[i] = nums[i]
// 		} else {
// 			leftMin[i] = mylibs.Min(leftMin[i-1], nums[i])
// 		}
// 	}
// 	for j := n - 1; j > -1; j-- {
// 		if j == n-1 {
// 			rightMax[j] = nums[j]
// 		} else {
// 			rightMax[j] = mylibs.Max(rightMax[j+1], nums[j])
// 		}
// 	}
// 	fmt.Println(leftMin)
// 	fmt.Println(rightMax)
// 	// 遍历
// 	for k := 1; k < n-1; k++ {
// 		if nums[k] > leftMin[k-1] && nums[k] < rightMax[k+1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// 解法二
// 贪心算法
// 思想: 为了找到递增的三元子序列，first\textit{first}first 和 second\textit{second}second 应该尽可能地小，此时找到递增的三元子序列的可能性更大。
func increasingTriplet(nums []int) bool {
	first, second := nums[0], math.MaxInt32
	if len(nums) > 2 {
		for _, num := range nums[1:] {
			if num <= first {
				first = num
			} else if num <= second {
				second = num
			} else { // first < second < num
				return true
			}
		}
	}
	return false
}

func IncreasingTriplet(nums []int) bool {
	return increasingTriplet(nums)
}
