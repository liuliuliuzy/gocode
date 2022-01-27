package solutions

import "sort"

//想到一个两次遍历的方法
//没看长度范围，超时了=。=
// func maxFrequency(nums []int, k int) int {
// 	qc(nums, 0, len(nums)-1)
// 	fmt.Println(nums)
// 	res := 1
// 	for i := len(nums) - 1; i > 0; i-- {
// 		m := k
// 		tmpIndex := i - 1
// 		tmpValue := 1
// 		for m > 0 && tmpIndex >= 0 {
// 			m -= (nums[i] - nums[tmpIndex])
// 			if m >= 0 {
// 				tmpValue++
// 				tmpIndex--
// 			}
// 		}
// 		fmt.Println(nums[i], tmpValue)
// 		res = max(res, tmpValue)
// 	}
// 	return res
// }

//排序+滑动窗口
func maxFrequency(nums []int, k int) int {
	//排序
	sort.Ints(nums)
	var (
		l     = 0 //窗口的左右边界
		r     = 1 //窗口的左右边界
		total = 0
		res   = 1 //结果最小也是1，所以res初始值定为1
	)
	for r < len(nums) {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= (nums[r] - nums[l])
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func MaxFrequency(nums []int, k int) int {
	return maxFrequency(nums, k)
}

// //快速排序
// func qc(nums []int, left, right int) {
// 	if left < right {
// 		i, j := left, right
// 		pivot := nums[i]
// 		for i < j {
// 			for i < j && nums[j] > pivot {
// 				j--
// 			}
// 			//找到一个比pivot小的，就将它放到当前空缺的位置
// 			nums[i] = nums[j]
// 			for i < j && nums[i] < pivot {
// 				i++
// 			}
// 			nums[j] = nums[i]
// 		}
// 		nums[i] = pivot
// 		qc(nums, left, i-1)
// 		qc(nums, i+1, right)
// 	}
// }
