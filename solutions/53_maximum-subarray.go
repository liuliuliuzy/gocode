package solutions

// 方法1：前缀和数组，然后暴力 O(n^2)
// 焯！评论里哪个b说暴力可以过的...
// func maxSubArrays(num []int) int {
// 	ans := math.MinInt32
// 	sum := make([]int, len(num)+1)
// 	for i, item := range num {
// 		for j := i + 1; j < len(num)+1; j++ {
// 			sum[j] += item
// 		}
// 	}
// 	// fmt.Println(sum)
// 	for i := 0; i < len(num); i++ {
// 		for j := i + 1; j < len(num)+1; j++ {
// 			if sum[j]-sum[i] > ans {
// 				ans = sum[j] - sum[i]
// 			}
// 		}
// 	}
// 	return ans
// }

// 方法2：动态规划
// f(i)=max{f(i−1)+nums[i],nums[i]}
func maxSubArrays(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

func MaxSubArrays(nums []int) int {
	return maxSubArrays(nums)
}
