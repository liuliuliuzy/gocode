package solutions

//二分查找（高阶方法）
// func lengthOfLIS(nums []int) int {
// 	s := []int{}
// 	for _, num := range nums {
// 		//二分查找
// 		i := sort.SearchInts(s, num)
// 		if i == len(s) {
// 			s = append(s, 0)
// 		}
// 		s[i] = num
// 	}
// 	fmt.Println(s)
// 	return len(s)
// }

//试试动态规划
func lengthOfLIS(nums []int) int {
	dp := []int{1}
	res := 1
	for i := 1; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func LengthOfLIS(nums []int) int {
	return lengthOfLIS(nums)
}
