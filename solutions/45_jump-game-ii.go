package solutions

//dp
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	//初始值
	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		tmp := n
		for j := 0; j < nums[i]; j++ {
			if i+1+j < n {
				// fmt.Println(i, j, tmp, dp[i+1+j])
				tmp = min(tmp, 1+dp[i+1+j])
			}
		}
		dp[i] = tmp
		// fmt.Printf("dp[%d] = %d\n", i, tmp)
	}
	return dp[0]
}

func Jump(nums []int) int {
	return jump(nums)
}
