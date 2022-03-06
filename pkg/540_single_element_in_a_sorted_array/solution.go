package singleelementinasortedarray

// 二分查找
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 这里抓住mid的奇偶性
		mid := (l + r) / 2
		// 偶数就是加1，奇数就是减1
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func SingleNonDuplicate(nums []int) int {
	return singleNonDuplicate(nums)
}
