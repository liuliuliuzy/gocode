package solutions

//题目要求 O(log n)，所以需要使用二分算法
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	ans := len(nums)
	//二分查找，找到有序数组nums中第一个≥target的元素索引
	for start <= end {
		mid := (end-start)>>1 + start
		if target <= nums[mid] {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
