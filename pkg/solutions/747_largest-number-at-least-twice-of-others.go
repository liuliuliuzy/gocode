package solutions

import "fmt"

/*
给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。

请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。
*/

// find out the top two number in the nums
func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	first, second := 0, 1
	if nums[second] > nums[first] {
		first, second = second, first
	}
	fmt.Println(first, second)
	for index := 2; index < len(nums); index++ {
		if nums[index] > nums[first] {
			// update first
			second = first
			first = index
		} else if nums[index] > nums[second] {
			second = index
		} else {
			continue
		}
	}
	fmt.Println(first, second)
	if nums[first] >= nums[second]*2 {
		return first
	} else {
		return -1
	}
}

func DominantIndex(nums []int) int {
	return dominantIndex(nums)
}
